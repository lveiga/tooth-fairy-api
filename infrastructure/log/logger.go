package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmlogrus"
)

// Logger - adapter to log mechanism
type Logger interface {
	GetLogger() *logrus.Logger
}

type loggerAdapter struct {
	logger *logrus.Logger
}

func (l *loggerAdapter) GetLogger() *logrus.Logger {
	return l.logger
}

// NewLogger - create logger instance
func NewLogger(environment string) Logger {
	var logger = &logrus.Logger{
		Out:   os.Stderr,
		Hooks: make(logrus.LevelHooks),
		Level: logrus.DebugLevel,
	}

	if environment != "prd" {
		logger.Formatter = &logrus.TextFormatter{}
		logger.Level = logrus.DebugLevel
	} else {
		logger.Formatter = &logrus.JSONFormatter{}
		logger.Level = logrus.TraceLevel
	}

	apm.DefaultTracer.SetLogger(logger)
	logrus.AddHook(&apmlogrus.Hook{})
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	return &loggerAdapter{
		logger: logger,
	}
}
