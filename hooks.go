package logger

import (
	"os"

	"github.com/nori-io/nori-common/logger"
)

// Internal type for storing the hooks on a logger instance.
type LevelHooks map[logger.Level][]logger.Hook

// Add a hook to an instance of logger. This is called with
// `log.Hooks.Add(new(MyHook))` where `MyHook` implements the `Hook` interface.
func (hooks LevelHooks) Add(hook logger.Hook) {
	for _, level := range hook.Levels() {
		hooks[level] = append(hooks[level], hook)
	}
}

// Fire all the hooks for the passed level. Used by `entry.log` to fire
// appropriate hooks for a log entry.
func (hooks LevelHooks) Fire(level logger.Level, message []byte) error {
	for _, hook := range hooks[level] {
		if err := hook.Fire(level, message); err != nil {
			return err
		}
	}
	return nil
}

type FileHook struct {
	Writer *os.File
}

func NewFileHook(name string) (*FileHook, error) {
	file, err := os.Create(name)
	if err == nil {
		return &FileHook{Writer: file}, err
	}
	return nil, err
}

func (hook *FileHook) Levels() []logger.Level {
	return []logger.Level{logger.LevelFatal, logger.LevelPanic, logger.LevelNotice, logger.LevelCritical, logger.LevelError,
		logger.LevelWarning, logger.LevelInfo, logger.LevelDebug}
}
func (hook *FileHook) Fire(level logger.Level, message []byte) error {

	switch level {
	case logger.LevelCritical:
		hook.Writer.Write(message)
	case logger.LevelDebug:
		hook.Writer.Write(message)
	case logger.LevelError:
		hook.Writer.Write(message)
	case logger.LevelFatal:
		hook.Writer.Write(message)
	case logger.LevelInfo:
		hook.Writer.Write(message)
	case logger.LevelNotice:
		hook.Writer.Write(message)
	case logger.LevelPanic:
		hook.Writer.Write(message)

	case logger.LevelWarning:
		hook.Writer.Write(message)

	default:
		return nil
	}
	return nil
}
