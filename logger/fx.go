package logger

import (
	"github.com/epes/econfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Error(args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatal(args ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Info(args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
}

func New(environment econfig.Environment) (Logger, error) {
	var logConfig zap.Config
	if environment == econfig.EnvironmentProduction {
		//logConfig = zap.NewProductionConfig()
		logConfig = zap.NewDevelopmentConfig()
	} else {
		logConfig = zap.NewDevelopmentConfig()
	}

	var logLevel zapcore.Level
	if environment == econfig.EnvironmentProduction {
		logLevel = zapcore.InfoLevel
	} else {
		logLevel = zapcore.DebugLevel
	}

	logConfig.Level.SetLevel(logLevel)

	logger, err := logConfig.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
