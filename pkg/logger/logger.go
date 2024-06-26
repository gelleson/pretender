package logger

import "go.uber.org/zap"

var logger *zap.Logger

func init() {
	logger, _ = zap.NewDevelopment()
}

func L() *zap.Logger {
	return logger
}

func Named(name string) *zap.Logger {
	return logger.Named(name)
}
