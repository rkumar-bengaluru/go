package log

import (
	"github.com/rkumar-bengaluru/go/logger/console"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Module = fx.Options(
	fx.Provide(
		NewDevelopmentLogger,
	),
)

func NewDevelopmentLogger() *console.ConsoleLogger {
	l := console.NewConsoleLogger(func(z *zap.Config) {
		z.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	})
	return l
}
