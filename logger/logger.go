package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//func NewLogger() {
//	logLevel, err := zapcore.ParseLevel(os.Getenv("LOG_LEVEL"))
//	if err != nil {
//		panic(fmt.Sprintf("Unknow log level: %s", logLevel))
//	}
//
//	var cfg zap.Config
//	if os.Getenv("LOG_ENV") == "development" {
//		cfg = zap.NewDevelopmentConfig()
//		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
//
//	} else {
//		cfg = zap.NewProductionConfig()
//	}
//	logger, err := cfg.Build()
//	if err != nil {
//		logger = zap.NewNop()
//	}
//	zap.ReplaceGlobals(logger)
//}
//
//func Close() {
//	defer zap.L().Sync()
//}

var logger *zap.Logger

// Setjsonencoder sets the logger code
func setJSONEncoder() zapcore.Encoder {
	EncoderConfig := zap.NewProductionEncoderConfig()
	return zapcore.NewConsoleEncoder(EncoderConfig)
}

// Setloggerwrite sets the logger to write to the file
func setLoggerWrite() zapcore.WriteSyncer {
	l := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,
		MaxBackups: 1,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	}
	return zapcore.AddSync(l)
}

// Initlogger initialization logger
func InitLogger() {
	core := zapcore.NewCore(setJSONEncoder(), setLoggerWrite(), zap.InfoLevel)
	logger = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}
