package logger

import (
	"frame/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	v2 "gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func init() {
	var writer zapcore.WriteSyncer

	cfg := conf.GetConfig()
	switch cfg.Logger.Target {
	case "console":
		writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	case "file":
		w := zapcore.AddSync(&v2.Logger{
			Filename:   cfg.Logger.Filename,
			MaxSize:    50, // single file max 10 MB
			MaxBackups: 0,  // keep all
			MaxAge:     30, // keep 30 days
		})
		writer = zapcore.NewMultiWriteSyncer(w)
	default:
		writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	}

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(newEncoderConfig()),
		writer,
		resolveLevel(cfg.Logger.Level),
	)

	Logger = zap.New(core, zap.AddCaller())
	Sugar = Logger.Sugar()
}

// 决定打印日志的等级
func resolveLevel(l string) zapcore.Level {
	switch l {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.DebugLevel
	}
}
