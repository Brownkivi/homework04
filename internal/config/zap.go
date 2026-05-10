package config

import (
	"BlodWeb/configs"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZap() {
	// 日志文件切割
	writer := &lumberjack.Logger{
		Filename:   "../logs/server.log", // 日志文件
		MaxSize:    10,                   // 单个文件最大 10MB
		MaxBackups: 100,                  // 最多保留 100 个文件
		MaxAge:     30,                   // 保留 30 天
		Compress:   true,                 // 压缩
	}

	// 日志格式配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	// 同时输出到：控制台 + 文件
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(writer),
			zap.InfoLevel,
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			zap.DebugLevel,
		),
	)

	// 创建 Logger
	configs.Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer configs.Logger.Sync()
}
