package file

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Module = fx.Options(
	fx.Provide(NewFileLogger),
)

type Option func(*zap.Config)

type Logger interface {
	Debug(msg string)
	Error(msg string, err error)
	Info(msg string)
	Warn(msg string)
}

type RotationLogger struct {
	logger *zap.SugaredLogger
}

func getRotationEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getRotationLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./rotate.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
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

func NewFileLogger(setters ...Option) *RotationLogger {
	writerSyncer := getRotationLogWriter()
	encoder := getRotationEncoder()
	//appName, _ := os.Executable()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer l.Sync()
	return &RotationLogger{l.Sugar()}
}

// Debug
func (z *RotationLogger) Debug(msg string) {
	z.logger.Debugw(msg)
}

// Error ...
func (z *RotationLogger) Error(msg string, err error) {
	keyAndValues := []interface{}{}
	keyAndValues = append(keyAndValues, 0, err)
	z.logger.Errorw(msg, keyAndValues)
}

// Info ...
func (z *RotationLogger) Info(msg string) {
	z.logger.Infow(msg)
}

// Warn ...
func (z *RotationLogger) Warn(msg string) {
	z.logger.Warnw(msg)
}
