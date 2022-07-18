package server

import (
	"context"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(NewServer),
)

func NewServer(lc fx.Lifecycle, router *mux.Router) {
	server := &http.Server{}
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				server.Addr = ":8080"
				listener, _ := net.Listen("tcp", server.Addr)
				server.Handler = router
				go func() {
					server.Serve(listener)
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				server.Shutdown(ctx)
				return nil
			},
		},
	)
}
