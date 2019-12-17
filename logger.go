package logger

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/nori-io/nori-common/logger"
)

type Core struct {
	Fields []logger.Field
	Time   time.Time
	Level  logger.Level
}

type Logger struct {
	Out       io.Writer
	Mu        *sync.Mutex
	Core      Core
	Formatter Formatter
	Hooks     LevelHooks
}

type LevelEnabler interface {
	Enabled(logger logger.Level) bool
}

func New(options ...Option) (logger logger.Logger) {
	log := &Logger{
		Out:       nil,
		Mu:        &sync.Mutex{},
		Core:      Core{},
		Formatter: nil,
		Hooks:     nil,
	}
	return log.WithOptions(options...)
}

// Fatal
func (log *Logger) Panic(format string, opts ...interface{}) {
	log.Log(logger.LevelPanic, format, opts...)

}

// Fatal logs a message with fatal level and exit with status set to 1
func (log *Logger) Fatal(format string, opts ...interface{}) {
	log.Log(logger.LevelFatal, format, opts...)

}

// Critical push to log entry with critical level
func (log *Logger) Critical(format string, opts ...interface{}) {
	log.Log(logger.LevelCritical, format, opts...)
}

// Error push to log entry with error level
func (log *Logger) Error(format string, opts ...interface{}) {
	log.Log(logger.LevelError, format, opts...)

}

// Warning push to log entry with warning level
func (log *Logger) Warning(format string, opts ...interface{}) {
	log.Log(logger.LevelWarning, format, opts...)

}

// Notice push to log entry with notice level
func (log *Logger) Notice(format string, opts ...interface{}) {
	log.Log(logger.LevelNotice, format, opts...)

}

// Info push to log entry with info level
func (log *Logger) Info(format string, opts ...interface{}) {
	log.Log(logger.LevelInfo, format, opts...)
}

// Debug push to log entry with debug level
func (log *Logger) Debug(format string, opts ...interface{}) {
	log.Log(logger.LevelDebug, format, opts...)
}

// Log push to log with specified level
func (log *Logger) Log(level logger.Level, format string, opts ...interface{}) {
	log.Mu.Lock()
	defer log.Mu.Unlock()

	//	log.Formatter.Format(log.Core.Fields...)

	for _, value := range log.Core.Fields {
		bytes, err := log.Formatter.Format(value)
		if err == nil {

			log.Out.Write(bytes)
		}
		//log.Out.Write([]byte(value.Key + " " + value.Value))
	}

	log.Out.Write([]byte(fmt.Sprintf(format, opts...)))

}

func (log *Logger) With(fields ...logger.Field) logger.Logger {
	if len(fields) == 0 {
		return log
	}

	With(log, fields...)
	l := log.clone()

	return l
}

func (log *Logger) clone() *Logger {
	copy := *log

	copy.Mu= log.Mu
	return &copy
}

func With(log *Logger, fields ...logger.Field) *Logger {

	clone := log
	tempCoreFields := log.Core.Fields
	clone.Core.Fields = append(clone.Core.Fields, fields...)
	log.Core.Fields = tempCoreFields

	return log
}

func (log *Logger) WithOptions(opts ...Option) *Logger {
	c := log.clone()
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}
