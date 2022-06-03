package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type RotationLogger struct {
	zap *zap.SugaredLogger
}

var appName string

func NewRotationLogger() *RotationLogger {
	writerSyncer := getRotationLogWriter()
	encoder := getRotationEncoder()
	//appName, _ := os.Executable()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	l := zap.New(core, zap.AddCaller())
	defer l.Sync()

	return &RotationLogger{l.Sugar()}
}

func (logger *RotationLogger) Debug(msg string) {
	logger.zap.Debug(msg)
}

func (logger *RotationLogger) Warn(msg string) {
	logger.zap.Warn(msg)
}

func (logger *RotationLogger) Error(msg string) {
	logger.zap.Error(msg)
}

func (logger *RotationLogger) Info(msg string) {
	logger.zap.Info(msg)
}

func getRotationEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getRotationLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./roatet.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
