package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rkumar-bengaluru/go/logger"
	"github.com/rkumar-bengaluru/go/rest/model"
)

var alog = logger.NewRotationLogger()

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, pwd, dbname string) {
	connectionString := fmt.Sprintf("port=49154 user=%v password=%v dbname=%v sslmode=disable",
		user, pwd, dbname)
	alog.Info(connectionString)
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
	}

	alog.Info(fmt.Sprintf("DB Initialized %v", a.DB))
	a.Router = mux.NewRouter()
	alog.Info(fmt.Sprintf("Router %v\n", a.Router))
	a.initRoutes()
}

// routes
func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	alog.Info(fmt.Sprintf("getProduct::product to fetch %v", id))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product id")
		return
	}
	p := model.Product{ID: id}
	if err = p.GetProduct(a.DB); err != nil {
		fmt.Println(err)
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
	var p model.Product
	alog.Info(fmt.Sprintf("recvd::%v", r.Body))
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if err := p.CreateProduct(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, p)
}

func (a *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	var p model.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if err := p.UpdateProduct(a.DB); err != nil {
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
	alog.Info(fmt.Sprintf("returing -> %v", string(response)))
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

	p := model.Product{ID: id}
	if err := p.DeleteProduct(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	alog.Info("call to getProducts...")
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := model.GetProducts(a.DB, start, count)
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
}

func (app App) Run(add string) {
	http.ListenAndServe(add, app.Router)
}
