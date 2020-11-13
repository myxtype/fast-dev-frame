package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	v2 "gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

const (
	level    = zap.InfoLevel    // 打印日志的等级
	target   = "console"        // 日志打印目标："file" or "console"
	filename = "log/stream.log" // 存放日志的文件
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
	switch target {
	case "console":
		writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	case "file":
		w := zapcore.AddSync(&v2.Logger{
			Filename:   filename,
			MaxSize:    10, // 10 MB
			MaxBackups: 0,  // keep all
			MaxAge:     7,  // keep 7 days
		})
		writer = zapcore.NewMultiWriteSyncer(w)
	}

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(newEncoderConfig()),
		writer,
		level,
	)

	Logger = zap.New(core, zap.AddCaller())
	Sugar = Logger.Sugar()
}
