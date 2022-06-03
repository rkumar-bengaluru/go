package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type JSONEncodeLogger struct {
	zap *zap.SugaredLogger
}

func NewJSONEncodeLogger() *JSONEncodeLogger {
	writerSyncer := getJsonLogWriter()
	encoder := getJsonEncoder()

	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	l := zap.New(core, zap.AddCaller())

	defer l.Sync()
	return &JSONEncodeLogger{l.Sugar()}
}

func (logger *JSONEncodeLogger) Debug(msg string) {
	logger.zap.Debug(msg)
}

func (logger *JSONEncodeLogger) Warn(msg string) {
	logger.zap.Warn(msg)
}

func (logger *JSONEncodeLogger) Error(msg string) {
	logger.zap.Error(msg)
}

func (logger *JSONEncodeLogger) Info(msg string) {
	logger.zap.Info(msg)
}

func getJsonEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getJsonLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./json.log")
	return zapcore.AddSync(file)
}
