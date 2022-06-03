package model

import (
	"database/sql"
	"fmt"

	"github.com/rkumar-bengaluru/go/logger"
)

var log = logger.NewRotationLogger()

type Product struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) GetProduct(db *sql.DB) error {
	log.Info(fmt.Sprintf("product to fetch %v\n", p.ID))
	return db.QueryRow("SELECT name,price from products where id=$1", p.ID).Scan(&p.Name, &p.Price)
}

func (p *Product) UpdateProduct(db *sql.DB) error {
	_, err := db.Exec("UPDATE products SET name=$1, price=$2 where id=$3", p.Name, p.Price, p.ID)
	return err
}

func (p *Product) DeleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE from products where id=$1", p.ID)
	return err
}

func (p *Product) CreateProduct(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT into products (name,price) VALUES($1,$2) RETURNING ID",
		p.Name,
		p.Price).Scan(&p.ID)
	return err
}

func GetProducts(db *sql.DB, start, count int) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price from products LIMIT $1 OFFSET $2", count, start)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	res := []Product{}
	fmt.Println(rows)
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		log.Info(fmt.Sprintf("id %v, name %v, price %v", p.ID, p.Name, p.Price))
		if err != nil {
			return nil, err
		}
		res = append(res, p)
	}
	return res, nil
}
