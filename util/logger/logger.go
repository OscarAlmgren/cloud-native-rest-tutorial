package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger *zerolog.Logger
}

func New(isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	return &Logger{logger: &logger}
}

func NewConsole(isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel((logLevel))

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &Logger{logger: &logger}
}

func (l *Logger) Output(w io.Writer) zerolog.Logger {
	return l.logger.Output(w)
}
