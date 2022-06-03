package main

import (
	"github.com/rkumar-bengaluru/go/logger"
	"github.com/rkumar-bengaluru/go/rest/routes"
)

func main() {
	logger.New().Info("default logger")
	logger.NewDevelopmentLogger().Info("from development logger")
	logger.NewRotationLogger().Info("from rotation logger")

	var a = routes.App{}
	a.Initialize(
		"postgres",
		"postgrespw",
		"sampledb",
	)

	a.Run(":8080")
}
