package logger

import (
	"io"

	logger "github.com/nori-io/logger/hooks/syslog"
)

// An Option configures a Logger.
type Option interface {
	apply(*Logger)
}

// optionFunc wraps a func so it satisfies the Option interface.
type optionFunc func(*Logger)

func (f optionFunc) apply(log *Logger) {
	f(log)
}

// ErrorOutput sets the destination for errors generated by the Logger. Note
// that this option only affects internal errors; for sample code that sends
// error-level logs to a different location from info- and debug-level logs,
// see the package-level AdvancedConfiguration example.
//
// The supplied WriteSyncer must be safe for concurrent use. The Open and
// zapcore.Lock functions are the simplest ways to protect files with a mutex.
func SetJsonFormatter() Option {
	return optionFunc(func(log *Logger) {
		log.Formatter = &JSONFormatter{
			TimestampFormat:  "",
			DisableTimestamp: false,
			DataKey:          "",
			FieldMap:         nil,
			PrettyPrint:      false,
		}
	})
}

func SetOutWriter(writer io.Writer) Option {
	return optionFunc(func(log *Logger) {
		log.Out = &writer
	})
}

func SetSysLogHook(hook logger.SyslogHook) Option {
	return optionFunc(func(log *Logger) {
		log.Hooks = hook
	})
}

/*// WrapCore wraps or replaces the Logger's underlying zapcore.Core.
func WrapCore(f func(zapcore.Core) zapcore.Core) Option {
	return optionFunc(func(log *Logger) {
		log.core = f(log.core)
	})
}

// Hooks registers functions which will be called each time the Logger writes
// out an Entry. Repeated use of Hooks is additive.
//
// Hooks are useful for simple side effects, like capturing metrics for the
// number of emitted logs. More complex side effects, including anything that
// requires access to the Entry's structured fields, should be implemented as
// a zapcore.Core instead. See zapcore.RegisterHooks for details.
func Hooks(hooks ...func(zapcore.Entry) error) Option {
	return optionFunc(func(log *Logger) {
		log.core = zapcore.RegisterHooks(log.core, hooks...)
	})
}

// Fields adds fields to the Logger.
func Fields(fs ...Field) Option {
	return optionFunc(func(log *Logger) {
		log.core = log.core.With(fs)
	})
}

// ErrorOutput sets the destination for errors generated by the Logger. Note
// that this option only affects internal errors; for sample code that sends
// error-level logs to a different location from info- and debug-level logs,
// see the package-level AdvancedConfiguration example.
//
// The supplied WriteSyncer must be safe for concurrent use. The Open and
// zapcore.Lock functions are the simplest ways to protect files with a mutex.
func ErrorOutput(w zapcore.WriteSyncer) Option {
	return optionFunc(func(log *Logger) {
		log.errorOutput = w
	})
}

// Development puts the logger in development mode, which makes DPanic-level
// logs panic instead of simply logging an error.
func Development() Option {
	return optionFunc(func(log *Logger) {
		log.development = true
	})
}

// AddCaller configures the Logger to annotate each message with the filename
// and line number of zap's caller.
func AddCaller() Option {
	return optionFunc(func(log *Logger) {
		log.addCaller = true
	})
}

// AddCallerSkip increases the number of callers skipped by caller annotation
// (as enabled by the AddCaller option). When building wrappers around the
// Logger and SugaredLogger, supplying this Option prevents zap from always
// reporting the wrapper code as the caller.
func AddCallerSkip(skip int) Option {
	return optionFunc(func(log *Logger) {
		log.callerSkip += skip
	})
}

// AddStacktrace configures the Logger to record a stack trace for all messages at
// or above a given level.
func AddStacktrace(lvl zapcore.LevelEnabler) Option {
	return optionFunc(func(log *Logger) {
		log.addStack = lvl
	})
}
*/
