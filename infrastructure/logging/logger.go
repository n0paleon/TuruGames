package logging

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SLogger = zap.SugaredLogger

func NewLogger(config *viper.Viper) *SLogger {
	options := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.Level(config.GetInt("log.level"))),
		Development: config.GetBool("app.development"),
		Encoding:    config.GetString("log.encoding"),
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := options.Build()
	if err != nil {
		panic(err)
	}

	return logger.Sugar()
}
