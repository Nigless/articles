package logger

import (
	"strings"

	"github.com/rs/zerolog"
)

func New(level string) zerolog.Logger {
	var l zerolog.Level
	level = strings.ToLower(level)
	switch level {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	out := zerolog.NewConsoleWriter()
	logger := zerolog.New(out).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return logger
}
