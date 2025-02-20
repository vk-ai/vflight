package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log  *zap.Logger
	once sync.Once
)

// InitLogger initializes the logger with the given environment
// env can be "development" or "production"
func InitLogger(env string) *zap.Logger {
	once.Do(func() {
		var config zap.Config

		if env == "production" {
			config = zap.NewProductionConfig()
		} else {
			config = zap.NewDevelopmentConfig()
		}

		// Configure the logger output
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		config.EncoderConfig = encoderConfig

		var err error
		log, err = config.Build(zap.AddCallerSkip(1))
		if err != nil {
			panic(err)
		}
	})

	return log
}

// GetLogger returns the singleton logger instance
func GetLogger() *zap.Logger {
	if log == nil {
		return InitLogger("development")
	}
	return log
}

// Info logs a message at info level
func Info(msg string, fields ...zap.Field) {
	GetLogger().Info(msg, fields...)
}

// Error logs a message at error level
func Error(msg string, fields ...zap.Field) {
	GetLogger().Error(msg, fields...)
}

// Debug logs a message at debug level
func Debug(msg string, fields ...zap.Field) {
	GetLogger().Debug(msg, fields...)
}

// Warn logs a message at warn level
func Warn(msg string, fields ...zap.Field) {
	GetLogger().Warn(msg, fields...)
}

// Fatal logs a message at fatal level and then calls os.Exit(1)
func Fatal(msg string, fields ...zap.Field) {
	GetLogger().Fatal(msg, fields...)
	os.Exit(1)
}
