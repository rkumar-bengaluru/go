package main

import (
	"github.com/rkumar-bengaluru/go/zaplog"
)

func main() {
	zaplog.Initialize()
	zaplog.Logger.Info("debug message")
	zaplog.InitializeSugaredLogger()
	zaplog.SugarLogger.Info("logging to sugared logger")

	zaplog.InitFileLogger()
	zaplog.FileLogger.Info("writing log to file")
}
