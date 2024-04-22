package logger

import (
	"go.uber.org/zap"
)

func Init(level string) (*zap.Logger, error) {
	switch level {
	case "dev":
		return zap.NewDevelopment()
	default:
		return zap.NewProduction()
	}
}
