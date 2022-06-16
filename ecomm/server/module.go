package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rkumar-bengaluru/go/ecomm/config"
	"github.com/rkumar-bengaluru/go/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewServer),
	fx.Invoke(InitServer),
)

type CatalogServerInterface interface {
	Start()
	Stop()
}

type CatalogServer struct {
	log    *logger.Logger
	config *config.EcommConfigReader
	Router *mux.Router
}

func (c *CatalogServer) Start() {
	port, _ := c.config.GetStringKey("server.port")
	c.log.Info(fmt.Sprintf("starting server on port %v\n", port))
	fmt.Println(c.Router)
	err := http.ListenAndServe(":"+port, c.Router)
	fmt.Println(err)
}

func (c *CatalogServer) Stop() {
	c.log.Info("Stopping Server")
}

func NewServer(
	lc fx.Lifecycle,
	log *logger.Logger,
	config *config.EcommConfigReader,
	r *mux.Router) *CatalogServer {
	server := &CatalogServer{
		log:    log,
		config: config,
		Router: r,
	}

	return server
}
func InitServer(s *CatalogServer) {
	s.log.Info("initializing server...")
}
