package model

import (
	"github.com/rkumar-bengaluru/go/config"
	"github.com/rkumar-bengaluru/go/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewCatalog),
)

func NewCatalog(log *logger.Logger, c *config.ConfigReader) ProductCatalog {
	log.Info("initializing catalog...")
	if k, _ := c.GetStringKey("db.type"); k == "DB" {
		log.Info("initializing db catalog")
		return &DBProductCatalog{log: log}
	} else {
		log.Info("initializing inmem catalog")
		return &InMemoryProductCatalog{log: log}
	}
}
