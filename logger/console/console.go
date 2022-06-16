package console

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Module = fx.Options(
	fx.Provide(NewConsoleLogger),
)

type Option func(*zap.Config)
type Logger interface {
	Debug(msg string)
	Error(msg string, err error)
	Info(msg string)
	Warn(msg string)
}

type ConsoleLogger struct {
	logger *zap.SugaredLogger
}

func NewZapConsoleLogger(setters ...Option) *zap.Logger {
	// Default config
	config := zap.Config{
		Encoding:          "json",
		Level:             zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableStacktrace: false,
		EncoderConfig:     NewConsoleZapEncoderConfig(),
	}

	for _, setter := range setters {
		setter(&config)
	}

	// We need to skip one caller, since we are going to wrap some functions
	logger, _ := config.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	zap.ReplaceGlobals(logger)

	return logger
}

func NewConsoleZapEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		CallerKey:     "caller",
		FunctionKey:   "func",
		StacktraceKey: "stacktrace",
		TimeKey:       "time",
		LevelKey:      "level",
		MessageKey:    "msg",
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime:    zapcore.RFC3339TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}
}

func NewConsoleLogger(setters ...Option) *ConsoleLogger {
	return &ConsoleLogger{NewZapConsoleLogger(setters...).Sugar()}
}

// Debug
func (z *ConsoleLogger) Debug(msg string) {
	z.logger.Debugw(msg)
}

// Error ...
func (z *ConsoleLogger) Error(msg string, err error) {
	keyAndValues := []interface{}{}
	keyAndValues = append(keyAndValues, 0, err)
	z.logger.Errorw(msg, keyAndValues)
}

// Info ...
func (z *ConsoleLogger) Info(msg string) {
	z.logger.Infow(msg)
}

// Warn ...
func (z *ConsoleLogger) Warn(msg string) {
	z.logger.Warnw(msg)
}
