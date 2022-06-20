package model

import (
	"database/sql"
	"errors"

	"fmt"

	_ "github.com/lib/pq"
	"github.com/rkumar-bengaluru/go/logger"
)

type DBProductCatalog struct {
	log *logger.Logger
	db  *sql.DB
}

func (c *DBProductCatalog) GetProduct(id int) (*Product, error) {
	p := Product{}
	c.log.Info(fmt.Sprintf("query db for product %d\n", id))
	rows, err := c.db.Query("select id,name,description,price from product where id=$1",
		id)

	if rows.Next() {
		c.log.Info("reading record...")
		rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
	}
	c.log.Info(fmt.Sprintf("product name %v\n", p.Name))

	if err != nil {
		return nil, errors.New("no product found...")
	}
	return &p, nil
}

func (c *DBProductCatalog) UpdateProduct(p Product) (*Product, error) {
	return nil, errors.New("no product found...")
}

func (c *DBProductCatalog) DeleteProduct(id int) (*Product, error) {
	return nil, errors.New("no product found...")
}

func (c *DBProductCatalog) CreateProduct(p Product) (*Product, error) {
	pr := Product{}
	err := c.db.QueryRow("insert into products (name,description, price) values($1,$2,$3)",
		p.ID, p.Description, p.Price).Scan(&pr.ID, &pr.Name, &pr.Description, &pr.Price)
	if err != nil {
		return nil, errors.New("error creating product...")
	}
	return &pr, nil
}

func (c *DBProductCatalog) GetAllProducts() ([]Product, error) {
	c.log.Info("getting all products")
	rows, err := c.db.Query("select id, name, description, price from product")
	if err != nil {
		return nil, errors.New("no product found...")
	}
	c.log.Info("getting all products")
	var all []Product
	for rows.Next() {
		p := Product{}
		rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		all = append(all, p)
	}
	c.log.Info("getting all products")
	return all, nil

}
