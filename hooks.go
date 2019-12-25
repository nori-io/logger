package logger

import (
	"fmt"
	"os"

	"github.com/nori-io/nori-common/logger"
	"github.com/sirupsen/logrus"
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
func (hooks LevelHooks) Fire(level logger.Level, log []logger.Field) error {
	for _, hook := range hooks[level] {
		if err := hook.Fire(log...); err != nil {
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
func (hook *FileHook) Fire(fields ...logger.Field) error {

	switch fields[0].Value {
	case logger.LevelPanic.String():
		hook.Writer.Write([]byte(fmt.Sprint(fields)))
	case logger.LevelFatal.String():
		hook.Writer.Write([]byte(fmt.Sprint(fields)))
	case logger.LevelError.String():
		hook.Writer.Write([]byte(fmt.Sprint(fields)))
	case logger.LevelWarning.String():
		hook.Writer.Write([]byte(fmt.Sprint(fields)))
	case logger.LevelInfo.String():
		hook.Writer.Write([]byte(fmt.Sprint(fields)))
	case logger.LevelDebug.String(), logrus.TraceLevel.String():
		hook.Writer.Write([]byte(fmt.Sprint(fields)))
	default:
		return nil
	}
	return nil
}
