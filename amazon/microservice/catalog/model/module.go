package model

import (
	"fmt"

	"github.com/rkumar-bengaluru/go/config"
	"github.com/rkumar-bengaluru/go/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewCatalog),
)

func NewCatalog(log *logger.Logger, c *config.ConfigReader) ProductCatalog {
	r, _ := c.GetStringKey("db.type")
	log.Info(fmt.Sprintf("initializing catalog...%v\n", r))
	if k, _ := c.GetStringKey("db.type"); k == "DB" {
		log.Info("initializing db catalog")
		return &DBProductCatalog{log: log}
	} else {
		log.Info("initializing inmem catalog")
		return &InMemoryProductCatalog{log: log}
	}
}
