package main

import (
	"context"

	"github.com/rkumar-bengaluru/go/ecomm/config"
	"github.com/rkumar-bengaluru/go/ecomm/log"
	"go.uber.org/fx"
)

func registerHooks(
	lifecycle fx.Lifecycle, config *config.EcommConfigReader,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				config.Start()
				return nil
			},
			OnStop: func(context.Context) error {
				config.Stop()
				return nil
			},
		},
	)
}

func main() {
	fx.New(
		log.Module,
		config.Module,
	).Run()

}
