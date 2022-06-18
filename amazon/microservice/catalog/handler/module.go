package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rkumar-bengaluru/go/amazon/microservice/catalog/model"
	"github.com/rkumar-bengaluru/go/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	model.Module,
	fx.Provide(NewCatalogHandler),
	fx.Invoke(registerRoutes),
)

type CatalogHandler struct {
	log     *logger.Logger
	catalog model.ProductCatalog
}

func NewCatalogHandler(log *logger.Logger, catalog model.ProductCatalog) *CatalogHandler {
	return &CatalogHandler{log: log, catalog: catalog}
}

func (h *CatalogHandler) getProduct(w http.ResponseWriter, r *http.Request) {
	h.log.Info("recvd request for GetProduct")
	p := mux.Vars(r)
	id, _ := strconv.Atoi(p["id"])
	product, _ := h.catalog.GetProduct(id)
	res, err := json.Marshal(product)
	if err != nil {
		h.log.Error("error in handler getProduct", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	h.log.Info("sent the success response....")
}

func (h *CatalogHandler) deleteProduct(w http.ResponseWriter, r *http.Request) {

}

func (h *CatalogHandler) updateProduct(w http.ResponseWriter, r *http.Request) {

}

func (h *CatalogHandler) createProduct(w http.ResponseWriter, r *http.Request) {

}

func (h *CatalogHandler) getAllProducts(w http.ResponseWriter, r *http.Request) {

}
func registerRoutes(router *mux.Router, h *CatalogHandler) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}).Methods("GET")
	// get
	router.HandleFunc("/product/{id}", h.getProduct).Methods("GET")
	// delete
	router.HandleFunc("/product/{id})", h.deleteProduct).Methods("DELETE")
	// update
	router.HandleFunc("/product/{id})", h.updateProduct).Methods("UPDATE")
	// create
	router.HandleFunc("/product", h.createProduct).Methods("POST")
	// get all
	router.HandleFunc("/product", h.getAllProducts).Methods("GET")
}
