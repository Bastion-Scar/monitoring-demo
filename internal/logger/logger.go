package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logger() *zap.Logger {
	cfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "message",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), os.Stdout, zapcore.DebugLevel)
	sampledCore := zapcore.NewSamplerWithOptions(core, time.Second, 20, 1)

	logger := zap.New(sampledCore, zap.AddStacktrace(zapcore.ErrorLevel))
	return logger
}
