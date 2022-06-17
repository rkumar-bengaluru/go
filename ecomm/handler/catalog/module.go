package catalog

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rkumar-bengaluru/go/ecomm/log"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewCatalogHandler,
	),
	fx.Invoke(Initialize),
)

type CatalogHandlerInterface interface {
	GetProduct()
}

type Product struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}
type CatalogHandler struct {
	log      *log.DevelopmentLogger
	products map[string]Product
	router   *mux.Router
}

func (c *CatalogHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]
	c.log.Info(fmt.Sprintf("recvd request for GetUser %v \n", id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	p := c.products[""+id+""]
	re, _ := json.Marshal(p)
	w.Write(re)

}

func respondWithJson(c *CatalogHandler, w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	c.log.Info(fmt.Sprintf("returing -> %v", string(response)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func NewCatalogHandler(log *log.DevelopmentLogger, router *mux.Router) *CatalogHandler {
	return &CatalogHandler{log: log, products: make(map[string]Product), router: router}
}

func Initialize(c *CatalogHandler) {
	c.products["1"] = Product{ID: "1", Name: "Orange"}
	c.products["2"] = Product{ID: "1", Name: "Apple"}

	c.router.HandleFunc("/product/{id}", c.GetProduct).Methods("GET")
	c.log.Info("initialize catalog endpoints...")
}
