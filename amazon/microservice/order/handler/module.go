package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rkumar-bengaluru/go/amazon/microservice/order/model"
	"go.uber.org/fx"
)

type OrderHandler struct {
	router *mux.Router
	model  model.IOrder
}

var Module = fx.Options(
	model.Module,
	fx.Provide(mux.NewRouter),
	fx.Provide(NewOrderHandler),
	fx.Invoke(RegisterRoutes),
)

func NewOrderHandler(model model.IOrder, router *mux.Router) *OrderHandler {
	return &OrderHandler{model: model, router: router}
}

func (handler *OrderHandler) GetOrder(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	p := mux.Vars(req)
	oid, _ := strconv.Atoi(p["id"])
	order := handler.model.GetOrder(oid)
	res, _ := json.Marshal(order)
	w.Write(res)
}

func RegisterRoutes(router *mux.Router, handler *OrderHandler) {
	router.HandleFunc("/order/{id}", handler.GetOrder).Methods("GET")
}
