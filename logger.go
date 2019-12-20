package logger

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/nori-io/nori-common/logger"
)

type Logger struct {
	Out       *io.Writer
	Mu        *sync.Mutex
	Fields    []logger.Field
	Formatter *JSONFormatter
	Hooks     *Hook
}
type LevelEnabler interface {
	Enabled(logger logger.Level) bool
}

func New(options ...Option) (loggerNew logger.Logger) {
	log := &Logger{
		Out:       nil,
		Mu:        &sync.Mutex{},
		Fields:    make([]logger.Field, 0),
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
	//fmt.Println("level is", level)
	//	log.Formatter.Format(log.Core.Fields...)
	levelType := fmt.Sprintf("%s", level)
	levelType = strings.ToUpper(levelType)
	(*log.Out).Write([]byte("[" + levelType + "]"))
	for _, value := range log.Fields {
		bytes, err := log.Formatter.Format(value)
		if err == nil {

			(*log.Out).Write(bytes)
		}
	}

	text, _ := log.Formatter.Format(logger.Field{
		Key:   "Msg",
		Value: format,
	})
	(*log.Out).Write([]byte(text))

}

func (log *Logger) With(fields ...logger.Field) logger.Logger {
	if len(fields) == 0 {
		return log
	}
	temp := log.Fields

	With(log, fields...)
	l := log.clone()

	log.Fields = temp

	return l
}

func (log *Logger) clone() *Logger {
	copy := *log
	return &copy
}

func With(log *Logger, fields ...logger.Field) *Logger {

	clone := log
	clone.Fields = append(clone.Fields, fields...)
	log = clone
	return log
}

func (log *Logger) WithOptions(opts ...Option) *Logger {
	c := log.clone()
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}
