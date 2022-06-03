package main

import (
	"github.com/rkumar-bengaluru/go/logger"
)

func main() {
	logger.New().Info("default logger")
	logger.NewDevelopmentLogger().Info("from development logger")
	logger.NewRotationLogger().Info("from rotation logger")
}
