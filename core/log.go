package core

import (
	"fmt"
	"github.com/zhaogongchengsi/starter-gin/config"
	"github.com/zhaogongchengsi/starter-gin/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

func CreateLogger(cfg *config.Config) (*zap.Logger, error) {
	conf := cfg.Zap
	if ok, _ := utils.DirPathExists(conf.Director); !ok {
		fmt.Printf("Create %s folder", conf.Director)
		_ = os.Mkdir(conf.Director, os.ModePerm)
	}

	rotate := NewFileRotatelog(&conf)

	writeSyncer, err := rotate.GetWriteSyncer(zapcore.DebugLevel.String())

	if err != nil {
		return nil, err
	}

	encoder := zapcore.NewJSONEncoder(CreateZapConfig(&conf))
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)

	logger.Sugar()

	return logger, nil
}

func CreateZapConfig(cfg *config.Zap) zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  cfg.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     CustomTimeEncoder(cfg),
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func CustomTimeEncoder(cfg *config.Zap) func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	return func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(cfg.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
	}
}

func TransportLevel(level string) zapcore.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
