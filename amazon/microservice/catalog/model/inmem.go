package model

import (
	"errors"
	"fmt"

	"github.com/rkumar-bengaluru/go/logger"
)

type InMemoryProductCatalog struct {
	log *logger.Logger
}

var allproducts = []Product{
	Product{ID: 1, Name: "Apple", Description: "green apple", Price: 100},
	Product{ID: 2, Name: "Banana", Description: "green Banana", Price: 80},
	Product{ID: 3, Name: "Orange", Description: "green Orange", Price: 90},
	Product{ID: 4, Name: "PineApple", Description: "green PineApple", Price: 10},
}

func (c *InMemoryProductCatalog) GetProduct(id int) (*Product, error) {
	c.log.Info(fmt.Sprintf("query on InMemoryProductCatalog for product %d\n", id))
	for i := 0; i < len(allproducts); i++ {
		if allproducts[i].ID == id {
			return &allproducts[i], nil
		}
	}
	return nil, errors.New("no product found...")
}
func (c *InMemoryProductCatalog) UpdateProduct(p Product) (*Product, error) {
	return nil, errors.New("no product found...")
}
func (c *InMemoryProductCatalog) DeleteProduct(id int) (*Product, error) {
	return nil, errors.New("no product found...")
}
func (c *InMemoryProductCatalog) CreateProduct(p Product) (*Product, error) {
	return nil, errors.New("no product found...")
}
func (c *InMemoryProductCatalog) GetAllProducts() ([]Product, error) {
	return nil, errors.New("no product found...")
}
