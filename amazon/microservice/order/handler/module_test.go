package handler

import (
	"github.com/gorilla/mux"
	"github.com/rkumar-bengaluru/go/amazon/microservice/order/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOrder(t *testing.T) {
	var inmem model.IOrder
	var router *mux.Router
	inmem = &model.InMemOrder{}
	router = mux.NewRouter()
	handler := &OrderHandler{router: router, model: inmem}
	RegisterRoutes(router, handler)
	req, _ := http.NewRequest("GET", "/order/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected %v got %v", http.StatusOK, rec.Code)
	}
}
