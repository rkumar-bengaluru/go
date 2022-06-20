package model

import (
	"database/sql"
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
	if k, _ := c.GetStringKey("db.type"); k == "postgres" {
		port, _ := c.GetStringKey("db.port")
		username, _ := c.GetStringKey("db.username")
		password, _ := c.GetStringKey("db.password")
		name, _ := c.GetStringKey("db.dbname")
		connectionString := fmt.Sprintf(
			"port=%v user=%v password=%v dbname=%v sslmode=disable",
			port,
			username,
			password,
			name)
		log.Info(connectionString)
		var err error
		db, err := sql.Open("postgres", connectionString)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		log.Info(fmt.Sprintf("DB Initialized %v", db))

		return &DBProductCatalog{log: log, db: db}
	} else {
		log.Info("initializing inmem catalog")
		return &InMemoryProductCatalog{log: log}
	}
}
