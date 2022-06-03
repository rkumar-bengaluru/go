package logger

import (
	"go.uber.org/zap"
)

type ProductionLogger struct {
	zap *zap.Logger
}

func New() *ProductionLogger {
	l, _ := zap.NewProduction()
	defer l.Sync()
	return &ProductionLogger{l}
}

func (logger *ProductionLogger) Debug(msg string) {
	logger.zap.Debug(msg)
}

func (logger *ProductionLogger) Warn(msg string) {
	logger.zap.Warn(msg)
}

func (logger *ProductionLogger) Error(msg string) {
	logger.zap.Error(msg)
}

func (logger *ProductionLogger) Info(msg string) {
	logger.zap.Info(msg)
}
