package bootstrap

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger(env string) *zap.Logger {
	var config zap.Config

	if env == "development" {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.ConsoleSeparator = " "

		config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("[15:04:05]:"))
		}

		// Disable stack traces in development for cleaner logs
		config.DisableStacktrace = true

	} else {
		// Use JSON encoding for production
		config = zap.NewProductionConfig()
		config.Encoding = "json"
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.LevelKey = "level"
		config.EncoderConfig.MessageKey = "message"
		config.EncoderConfig.CallerKey = "caller"
		config.EncoderConfig.StacktraceKey = "stacktrace"
	}

	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	config.EncoderConfig.LineEnding = "\n"

	logger := zap.Must(config.Build())

	return logger
}
