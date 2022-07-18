package main

import (
	"github.com/rkumar-bengaluru/go/amazon/microservice/order/handler"
	"github.com/rkumar-bengaluru/go/amazon/microservice/order/server"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		handler.Module,
		server.Module,
	)
	if app.Err() != nil {
		panic(app.Err())
	}
	app.Run()
}
