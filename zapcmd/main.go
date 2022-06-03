package main

import (
	"github.com/rkumar-bengaluru/go/logger"
)

func main() {
	logger.New().Info("info")
	logger.NewFileLogger().Info("info")
	logger.NewJSONEncodeLogger().Info("info")
	logger.NewRotationLogger().Info("info")
}
