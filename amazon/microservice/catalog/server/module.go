package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rkumar-bengaluru/go/config"
	"github.com/rkumar-bengaluru/go/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(mux.NewRouter),
	fx.Invoke(StartCatalogServer),
)

func StartCatalogServer(
	lc fx.Lifecycle,
	router *mux.Router,
	log *logger.Logger,
	aconfig *config.ConfigReader,
) {
	server := &http.Server{}
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Info("starting catalog amazon server....")
				port := aconfig.GetIntegerKey("server:port")
				idleTimeout := aconfig.GetIntegerKey("server:idle-timeout")
				readTimeout := aconfig.GetIntegerKey("server:read-timeout")
				writeTimeout := aconfig.GetIntegerKey("server:write-timeout")

				server.Addr = fmt.Sprintf(":%d", port)
				server.IdleTimeout = time.Duration(idleTimeout * time.Second)
				server.ReadTimeout = time.Duration(readTimeout * time.Second)
				server.WriteTimeout = time.Duration(writeTimeout * time.Second)
				listener, err := net.Listen("tcp", server.Addr)
				if err != nil {
					log.Error("error starting server", err)
					panic(err)
				}
				server.Handler = router
				go func() {
					server.Serve(listener)
				}()
				return nil
			},
			OnStop: func(c context.Context) error {
				log.Info("stopping the server")
				return server.Shutdown(c)
			},
		},
	)
}
