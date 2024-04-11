package logger

import (
	"github.com/palle-404/erp-be/src/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var logger *zap.Logger

func Log() *zap.Logger {
	return logger
}

func Init() (err error) {
	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	if config.AppCfg().GetString("log.level") == "debug" {
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	cfg := zap.Config{
		Encoding:         "json",
		Level:            level,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeTime:    customTimeEncoder, // Use custom time encoder
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}
	logger, err = cfg.Build()
	if err != nil {
		return err
	}
	logger.Info("Logger initialized ...")
	return nil
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("02-01-2006 15:04:05 IST"))
}
