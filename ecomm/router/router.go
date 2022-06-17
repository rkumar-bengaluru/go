package router

import (
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewRouter,
	),
)

func NewRouter() *mux.Router {
	return mux.NewRouter()
}
