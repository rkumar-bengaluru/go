package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

// Module ..
var Module = fx.Options(
	fx.Provide(
		mux.NewRouter,
	),
	fx.Invoke(
		StartServer,
	),
)

func StartServer(
	lc fx.Lifecycle,
	router *mux.Router,
) {
	server := &http.Server{}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				readTimeout := time.Duration(configServer.ReadTimeout) * time.Second
				writeTimeout := time.Duration(configServer.WriteTimeout) * time.Second
				idleTimeout := time.Duration(configServer.IdleTimeout) * time.Second

				server.Addr = fmt.Sprintf(":%d", configServer.Port)
				server.ReadTimeout = readTimeout
				server.WriteTimeout = writeTimeout
				server.IdleTimeout = idleTimeout
				server.Handler = router

				logger.Info("Starting the server...", log.Values{
					"server.addr": server.Addr,
				})

				listener, err := net.Listen("tcp", server.Addr)
				if err != nil {
					logger.Error("cannot start server", err)

					panic(err)
				}

				go func() {
					_ = server.Serve(listener)
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				logger.Info("Stopping the server...")

				return server.Shutdown(ctx)
			},
		},
	)
}
