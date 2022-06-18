package model

import (
	"errors"
	"fmt"

	"github.com/rkumar-bengaluru/go/logger"
)

type DBProductCatalog struct {
	log *logger.Logger
}

var alldbproducts []Product = []Product{
	Product{ID: 1, Name: "Apple", Description: "green apple", Price: 100},
	Product{ID: 2, Name: "Banana", Description: "green Banana", Price: 80},
	Product{ID: 3, Name: "Orange", Description: "green Orange", Price: 90},
	Product{ID: 4, Name: "PineApple", Description: "green PineApple", Price: 10},
}

func (c *DBProductCatalog) GetProduct(id int) (*Product, error) {
	c.log.Info(fmt.Sprintf("query on DBProductCatalog for product %d\n", id))
	for i := 0; i < len(alldbproducts); i++ {
		if alldbproducts[i].ID == id {
			return &alldbproducts[i], nil
		}
	}
	return nil, errors.New("no product found...")
}
func (c *DBProductCatalog) UpdateProduct(p Product) (*Product, error) {
	return nil, errors.New("no product found...")
}
func (c *DBProductCatalog) DeleteProduct(id int) (*Product, error) {
	return nil, errors.New("no product found...")
}
func (c *DBProductCatalog) CreateProduct(p Product) (*Product, error) {
	return nil, errors.New("no product found...")
}
func (c *DBProductCatalog) GetAllProducts() ([]Product, error) {
	return nil, errors.New("no product found...")
}
