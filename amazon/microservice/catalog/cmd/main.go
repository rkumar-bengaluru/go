package main

import (
	"github.com/rkumar-bengaluru/go/amazon/microservice/catalog/server"
	"github.com/rkumar-bengaluru/go/config"
	"github.com/rkumar-bengaluru/go/logger"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		logger.Module,
		config.Module,
		server.Module,
	)

	if app.Err() != nil {
		panic(app.Err())
	}

	app.Run()
}
