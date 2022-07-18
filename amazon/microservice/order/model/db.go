package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	c *sql.DB
}

func (db *PostgresDB) CreateOrder(pid, cid int) Order {
	return Order{}
}

func (p *PostgresDB) GetOrder(id int) Order {
	rows, err := p.c.Query("select id,pid,cid from product_order where id=$1", id)
	if err != nil {
		fmt.Errorf("Error %v", err)
	}
	order := Order{}
	if rows.Next() {
		rows.Scan(&order.ID, &order.PID, &order.CID)
	}
	return order
}
