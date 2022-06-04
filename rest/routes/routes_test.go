package routes_test

import (
	"github.com/rkumar-bengaluru/go/rest/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var app = routes.App{}

func TestGetProduct(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/product/1", nil)
	rec := httptest.NewRecorder()
	app.Router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected %v, recvd %v", http.StatusOK, rec.Code)
	}
}

func TestMain(m *testing.M) {
	app.Initialize(
		"postgres",
		"postgrespw",
		"sampledb",
	)
	ensureTableExist()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

var tableCreateQuery = `CREATE TABLE IF NOT EXISTS products (
		id serial primary key,
		name TEXT NOT NULL,
		price Numeric(10,2) NOT NULL DEFAULT 0.00
	)`

func ensureTableExist() {
	if err, _ := app.DB.Exec(tableCreateQuery); err != nil {
		log.Panic("db check failed...")
	}
}

func clearTable() {
	if err, _ := app.DB.Exec("delete from products"); err != nil {
		log.Panic("could not clear products table")
	}
}
