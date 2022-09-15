package utils

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func InitializeLogger(logFilename string) {

	core := zapcore.NewTee(
		zapcore.NewCore(getConsoleEncoder(), zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(getJSONEncoder(), getLogWriter(logFilename), zapcore.DebugLevel),
	)
	Logger = zap.New(core, zap.AddCaller()).Sugar()
	defer Logger.Sync()
}

func getEncoderConfig() zapcore.EncoderConfig {
	baseConfig := zap.NewProductionEncoderConfig()
	baseConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	baseConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	baseConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return baseConfig
}

func getJSONEncoder() zapcore.Encoder {
	config := getEncoderConfig()
	return zapcore.NewJSONEncoder(config)
}

func getConsoleEncoder() zapcore.Encoder {
	config := getEncoderConfig()
	return zapcore.NewConsoleEncoder(config)
}

func getLogWriter(logFilename string) zapcore.WriteSyncer {
	logsPath := "logs"
	if _, err := os.Stat(logsPath); os.IsNotExist(err) {
		os.Mkdir(logsPath, os.ModePerm)
	}
	file, _ := os.OpenFile(filepath.Join(logsPath, logFilename), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	return zapcore.AddSync(file)
}
