package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, pwd, dbname string) {
	connectionString := fmt.Sprintf("port=49154 user=%v password=%v dbname=%v sslmode=disable",
		user, pwd, dbname)
	fmt.Println(connectionString)
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Info(err)
	}
	fmt.Printf("DB Initialized %v", a.DB)
	a.Router = mux.NewRouter()
	fmt.Printf("Router %v\n", a.Router)
	a.initRoutes()
}

// routes
func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	log.Info("getProduct::product to fetch ", id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product id")
		return
	}
	p := Product{ID: id}
	if err = p.getProduct(a.DB); err != nil {
		log.Error(err)
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
			return
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	respondWithJson(w, http.StatusOK, p)
}

func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	log.Info("recvd::", r.Body)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		log.Error(err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if err := p.createProduct(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, p)
}

func (a *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if err := p.updateProduct(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, p)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	log.Info("returing ->", string(response))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	p := Product{ID: id}
	if err := p.deleteProduct(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	log.Info("call to getProducts...")
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := getProducts(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, products)
}

func addProduct(n int, a *App) {
	if n < 1 {
		n = 1
	}
	for i := 0; i < n; i++ {
		fmt.Println("inserted...")
		a.DB.Exec("insert into products(name,price) values ($1,$2", "P"+strconv.Itoa(i), (i+1.0)*10)
	}

}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.getProduct).Methods("GET")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
	fmt.Println("routes configured...")
	//addProduct(10, a)
	// p := Product{Name: "P1", Price: 66.77}
	// p.createProduct(a.DB)
}

func (app App) Run(add string) {
	log.Fatal(http.ListenAndServe(add, app.Router))
}
