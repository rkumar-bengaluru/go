package routes_test

import (
	"github.com/rkumar-bengaluru/go/restwgin/nroutes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProduct(t *testing.T) {
	app := routes.App{}
	app.InitializeRoutes()
	req, _ := http.NewRequest("GET", "/product/1", nil)
	rec := httptest.NewRecorder()
	app.Router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected 200 got %v", rec.Code)
	}
}

func TestGetProductNotAvailable(t *testing.T) {
	app := routes.App{}
	app.InitializeRoutes()
	req, _ := http.NewRequest("GET", "/product/10", nil)
	rec := httptest.NewRecorder()
	app.Router.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("expected 404 got %v", rec.Code)
	}
}
