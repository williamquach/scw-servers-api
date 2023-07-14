package utils

import (
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(),
	}
}

func (l *Logger) SetLogLevel(level string) {
	switch level {
	case "debug":
		l.SetLogLevel(log.DebugLevel.String())
	case "info":
		l.SetLogLevel(log.InfoLevel.String())
	case "warn":
		l.SetLogLevel(log.WarnLevel.String())
	case "error":
		l.SetLogLevel(log.ErrorLevel.String())
	default:
		l.SetLogLevel(log.InfoLevel.String())
	}
}

func (l *Logger) SetLogFormat(format string) {
	switch format {
	case "json":
		l.SetFormatter(&log.JSONFormatter{})
	case "text":
		l.SetFormatter(&log.TextFormatter{})
	default:
		l.SetFormatter(&log.TextFormatter{})
	}
}
