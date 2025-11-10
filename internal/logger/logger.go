package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	// Production config: JSON logs, fast, structured
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"} // logs to stdout (Docker-friendly)

	var err error
	Logger, err = cfg.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
}

// Sugar provides a sugared logger for easier syntax
func Sugar() *zap.SugaredLogger {
	return Logger.Sugar()
}
