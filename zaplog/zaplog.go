package zaplog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// https://sunitc.dev/2019/05/27/adding-uber-go-zap-logger-to-golang-project/
var Logger *zap.Logger

func Initialize() {
	Logger, _ = zap.NewProduction()
	defer Logger.Sync()
}

var SugarLogger *zap.SugaredLogger

func InitializeSugaredLogger() {
	l, _ := zap.NewProduction()
	SugarLogger = l.Sugar()
}

var FileLogger *zap.SugaredLogger

func InitFileLogger() {
	file, _ := os.Create("./test.log")
	writerSyncer := zapcore.AddSync(file)
	encoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	logger := zap.New(core)
	FileLogger = logger.Sugar()
}
