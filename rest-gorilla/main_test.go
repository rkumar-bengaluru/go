package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	log "github.com/sirupsen/logrus"
	"rupak.com/tmp"
)

const tableCreateQuery = `CREATE TABLE IF NOT EXISTS products
	(
		id serial PRIMARY KEY,
		name TEXT NOT NULL,
		price NUMERIC(10,2) NOT NULL DEFAULT 0.00
	);
`

var a = tmp.App{}

var q1 = `CREATE TABLE IF NOT EXISTS accounts (
	user_id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL,
        last_login TIMESTAMP
);`

func ensureTableExists() {

	if _, err := a.DB.Exec(tableCreateQuery); err != nil {
		log.Fatal(err)
	}
	fmt.Println("table created...")
}
func clearTable() {
	_, err := a.DB.Exec("delete from products;")
	if err != nil {
		log.Fatal("could not clear tables...")
	}
	a.DB.Exec("alter sequence products_id_seq restart with 1")
}

func TestMain(m *testing.M) {

	a.Initialize(
		"postgres",
		"postgrespw",
		"sampledb",
	)
	fmt.Printf("DB Initialized %v\n", a.Router)
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func executeRequest(r *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, r)
	return rr
}

func checkResponseCode(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("expected %v recieved %v", expected, actual)
	}
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/products", nil)
	rr := executeRequest(req)
	checkResponseCode(t, http.StatusOK, rr.Code)

	if body := rr.Body.String(); body != "[]" {
		t.Errorf("expected an empty array, got %s", body)
	}
}

func TestNonExistingProduct(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/product/11", nil)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, res.Code)
	var jsonR = make(map[string]string)
	json.Unmarshal(res.Body.Bytes(), &jsonR)
	log.Info(jsonR["error"])
	if jsonR["error"] != "Product not found" {
		t.Errorf("Expected the error key in reponse to be Product Not Found however recvd %v", jsonR["error"])
	}
}

func createProduct(t *testing.T) {
	var jsonProduct = []byte(`{"name": "P1", "price":66.67}`)

	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonProduct))
	res := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, res.Code)
	var jsonResp map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &jsonResp)
	if jsonResp["name"] != "P1" {
		t.Errorf("Expected P1 as product name, but recvd %v", jsonResp["name"])
	}

	if jsonResp["price"] != 66.67 {
		t.Errorf("Expected 66.77 as product price, but recvd %v", jsonResp["price"])
	}

	if jsonResp["id"] != 1.0 {
		t.Errorf("Expected product ID to be '1'. Got '%v'", jsonResp["id"])
	}
}
func TestCreateProduct(t *testing.T) {
	clearTable()
	createProduct(t)
}

func addProduct(n int) {
	if n < 1 {
		n = 1
	}
	for i := 0; i < n; i++ {
		a.DB.Exec("insert into products(name,price) values ($1,$2", "P"+strconv.Itoa(i), (i+1.0)*10)
	}
}

func TestGetProduct(t *testing.T) {
	clearTable()
	createProduct(t)
	req, _ := http.NewRequest("GET", "/product/1", nil)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)
}

func TestUpdateProduct(t *testing.T) {
	clearTable()
	createProduct(t)

	req, _ := http.NewRequest("GET", "/product/1", nil)
	response := executeRequest(req)
	var originalProduct map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalProduct)

	var jsonStr = []byte(`{"name":"test product - updated name", "price": 11.22}`)
	req, _ = http.NewRequest("PUT", "/product/1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalProduct["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalProduct["id"], m["id"])
	}

	if m["name"] == originalProduct["name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalProduct["name"], m["name"], m["name"])
	}

	if m["price"] == originalProduct["price"] {
		t.Errorf("Expected the price to change from '%v' to '%v'. Got '%v'", originalProduct["price"], m["price"], m["price"])
	}
}

func TestDeleteProduct(t *testing.T) {
	clearTable()
	createProduct(t)

	req, _ := http.NewRequest("GET", "/product/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/product/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/product/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}
