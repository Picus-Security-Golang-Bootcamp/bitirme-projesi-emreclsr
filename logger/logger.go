package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewLogger() {
	logLevel, err := zapcore.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		panic(fmt.Sprintf("Unknow log level: %s", logLevel))
	}

	var cfg zap.Config
	if os.Getenv("LOG_ENV") == "development" {
		cfg = zap.NewDevelopmentConfig()
		cfg.Level.SetLevel(logLevel)
		cfg.Encoding = "json"
		cfg.OutputPaths = []string{"zap.log"}
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	} else {
		cfg = zap.NewProductionConfig()
	}
	logger, err := cfg.Build()
	if err != nil {
		logger = zap.NewNop()
	}
	zap.ReplaceGlobals(logger)
}

func Close() {
	defer zap.L().Sync()
}
