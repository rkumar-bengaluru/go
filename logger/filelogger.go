package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FileLogger struct {
	zap *zap.SugaredLogger
}

func NewFileLogger() *FileLogger {
	writerSyncer := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	l := zap.New(core)

	defer l.Sync()
	return &FileLogger{l.Sugar()}
}

func (logger *FileLogger) Debug(msg string) {
	logger.zap.Debug(msg)
}

func (logger *FileLogger) Warn(msg string) {
	logger.zap.Warn(msg)
}

func (logger *FileLogger) Error(msg string) {
	logger.zap.Error(msg)
}

func (logger *FileLogger) Info(msg string) {
	logger.zap.Info(msg)
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}
