package logger

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/nori-io/nori-common/logger"
)

type Core interface {
	//LevelEnabler

	// With adds structured context to the Core.
	With([]Field) Core
	/*	// Check determines whether the supplied Entry should be logged (using the
		// embedded LevelEnabler and possibly some extra logic). If the entry
		// should be logged, the Core adds itself to the CheckedEntry and returns
		// the result.
		//
		// Callers must use Check before calling Write.
		Check(Entry, *CheckedEntry) *CheckedEntry
		// Write serializes the Entry and any Fields supplied at the log site and
		// writes them to their destination.
		//
		// If called, Write should always log the Entry and Fields; it should not
		// replicate the logic of Check.
		Write(Entry, []Field) error
		// Sync flushes buffered logs (if any).
		Sync() error*/
}
type WriteSyncer interface {
	io.Writer
	Sync() error
}

type Field struct {
	Key   string
	Value interface{}
}

type FieldType uint8

type Logger struct {
	Out  io.Writer
	mu   sync.Mutex
	core Core
}

var m logger.Fields

type ioCore struct {
	LevelEnabler
	key   string
	value interface{}
}
type LevelEnabler interface {
	Enabled(logger logger.Level) bool
}

func New() (logger logger.FieldLogger) {
	return Logger{Out: os.Stderr, mu: sync.Mutex{}}

}

// Fatal
func (log Logger) Panic(format string, opts ...interface{}) {
	log.Log(logger.LevelPanic, format, opts...)

}

// Fatal logs a message with fatal level and exit with status set to 1
func (log Logger) Fatal(format string, opts ...interface{}) {
	log.Log(logger.LevelFatal, format, opts...)

}

// Critical push to log entry with critical level
func (log Logger) Critical(format string, opts ...interface{}) {
	log.Log(logger.LevelCritical, format, opts...)
}

// Error push to log entry with error level
func (log Logger) Error(format string, opts ...interface{}) {
	log.Log(logger.LevelError, format, opts...)

}

// Warning push to log entry with warning level
func (log Logger) Warning(format string, opts ...interface{}) {
	log.Log(logger.LevelWarning, format, opts...)

}

// Notice push to log entry with notice level
func (log Logger) Notice(format string, opts ...interface{}) {
	log.Log(logger.LevelNotice, format, opts...)

}

// Info push to log entry with info level
func (log Logger) Info(format string, opts ...interface{}) {
	log.Log(logger.LevelInfo, format, opts...)
}

// Debug push to log entry with debug level
func (log Logger) Debug(format string, opts ...interface{}) {
	log.Log(logger.LevelDebug, format, opts...)
}

// Printf is like fmt.Printf, push to log entry with debug level
func (log Logger) Printf(format string, opts ...interface{}) {
	log.mu.Lock()
	defer log.mu.Unlock()
	log.Out.Write([]byte(fmt.Sprintf(format, opts...)))
}

// Write push to log entry with debug level
func (log Logger) Write(p []byte) (n int, err error) {
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.Out.Write(p)
}

// Log push to log with specified level
func (log Logger) Log(level logger.Level, format string, opts ...interface{}) {
	log.mu.Lock()
	defer log.mu.Unlock()
	log.Out.Write([]byte(fmt.Sprintf(format, opts...)))

}

func (log *Logger) WithField(key string, value interface{}) *Logger {
	if len(key) == 0 {
		return log
	}

	l := log.clone()
	l.Out = l.WithField(key, value)
	return l

	l = log.clone()

	return log
}

func (log *Logger) WithFields(fields logger.Fields) *Logger {
	if len(fields) == 0 {
		return log
	}
	fmt.Print(len(fields))

	l := log.clone()
	l.Out = l.WithFields(fields)

	return l
}

func (log *Logger) clone() *Logger {
	copy := *log
	return &copy
}

func (c *ioCore) With(fields []Field) Core {
	clone := c
	for i := 1; i < len(fields); i++ {
		clone.key = fields[i].Key
		clone.value = fields[i].Value
	}

	return clone
}
