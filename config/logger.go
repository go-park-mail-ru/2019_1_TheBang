package config

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func createGlobalLogger() *zap.SugaredLogger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		log.Fatalln("createLogger", err.Error)
	}

	// func ErrorOutput - определить куда летят внутренние ошибки логгера
	// toDo сделать advance cfg c определением места логирования

	return logger.Sugar()
}