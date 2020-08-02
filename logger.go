package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/nori-io/common/v3/logger"
	"github.com/nori-io/logger/internal/types"
)

type Logger struct {
	Out       io.Writer
	Mu        *sync.Mutex
	Fields    []logger.Field
	Formatter logger.Formatter
	Hooks     *types.LevelHooks
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
		Hooks:     &types.LevelHooks{},
	}

	log = log.WithOptions(options...)

	opts := []Option{}
	if log.Out == nil {
		opts = append(opts, SetOutWriter(os.Stderr))
	}
	if log.Formatter == nil {
		opts = append(opts, SetJsonFormatter(""))
	}
	if len(opts) > 0 {
		return log.WithOptions(opts...)
	}

	return log
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
	entry := logger.Entry{
		Formatter: log.Formatter,
		Level:     level,
		Time:      time.Now(),
		Message:   fmt.Sprintf(format, opts...),
	}
	// format output
	text, _ := log.Formatter.Format(entry, log.Fields...)

	// output
	log.Mu.Lock()
	defer log.Mu.Unlock()

	log.Out.Write(text)

	// fire hooks
	log.Hooks.Fire(entry, log.Fields...)
}

func (log *Logger) With(fields ...logger.Field) logger.Logger {
	if len(fields) == 0 {
		return log
	}
	l := log.clone()
	l.Fields = append(log.Fields, fields...)
	return l
}

func (log *Logger) WithOptions(opts ...Option) *Logger {
	c := log.clone()
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}

func (log *Logger) AddHook(hook logger.Hook) {
	log.Hooks.Add(hook)
}

func (log *Logger) clone() *Logger {
	copy := *log
	return &copy
}
