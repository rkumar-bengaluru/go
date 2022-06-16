package log

import (
	"github.com/rkumar-bengaluru/go/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Module = fx.Options(
	fx.Provide(
		NewDevelopmentLogger,
	),
)

func NewDevelopmentLogger() *logger.Logger {
	return logger.NewLogger(func(z *zap.Config) {
		z.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	})
}
