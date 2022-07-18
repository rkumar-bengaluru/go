package model

import (
	"database/sql"
	"fmt"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewModel),
)

func NewModel() IOrder {
	var m IOrder
	connectionString := fmt.Sprintf(
		"port=%v user=%v password=%v dbname=%v sslmode=disable",
		5432,
		"postgres",
		"password",
		"ecomm")
	fmt.Println(connectionString)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
	}
	m = &PostgresDB{c: db}
	return m
}
