package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rkumar-bengaluru/go/ecomm/config"
	"github.com/rkumar-bengaluru/go/ecomm/handler/catalog"
	"github.com/rkumar-bengaluru/go/ecomm/handler/user"
	"github.com/rkumar-bengaluru/go/ecomm/log"
	"github.com/rkumar-bengaluru/go/ecomm/router"
	"github.com/rkumar-bengaluru/go/ecomm/server"
	"go.uber.org/fx"
)

func registerHooks(
	lifecycle fx.Lifecycle, server *server.CatalogServer,
) {

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go server.Start()
				return nil
			},
			OnStop: func(context.Context) error {
				server.Stop()
				return nil
			},
		},
	)
}

func main() {
	app := fx.New(
		log.Module,
		config.Module,
		router.Module,
		catalog.Module,
		user.Module,
		server.Module,
		fx.Invoke(registerHooks),
	)
	app.Run()
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		fmt.Errorf("error %v\n", err)
	}

	if res, err := http.Get("http://localhost:8000/"); err != nil {
		fmt.Errorf("error %v\n", err)
	} else {
		fmt.Printf("got response %v\n", res)
	}

	stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		fmt.Errorf("error %v\n", err)
	}
}
