package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Cfg config of logger
var Cfg = zap.Config{
	Encoding:         "json",
	Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
	OutputPaths:      []string{"stdout"},
	ErrorOutputPaths: []string{"stderr"},
	EncoderConfig: zapcore.EncoderConfig{
		MessageKey: "message",
		LevelKey:   "level",
		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,
	},
}

//ErrorCfg config of error logger
var ErrorCfg = zap.Config{
	Encoding:    "json",
	Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
	OutputPaths: []string{"stderr"},
	EncoderConfig: zapcore.EncoderConfig{
		MessageKey: "message",
		LevelKey:   "level",
		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,
	},
}
