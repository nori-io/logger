package logger

import (
	"os"

	"github.com/nori-io/nori-common/logger"
)

// A hook to be fired when logging on the logging levels returned from
// `Levels()` on your implementation of the interface. Note that this is not
// fired in a goroutine or a channel with workers, you should handle such
// functionality yourself if your call is non-blocking and you don't wish for
// the logging calls for levels returned from `Levels()` to block.
type Hook interface {
	Levels() []logger.Level
	Fire(field []logger.Field) error
}

// Internal type for storing the hooks on a logger instance.
type LevelHooks map[logger.Level][]Hook

// Add a hook to an instance of logger. This is called with
// `log.Hooks.Add(new(MyHook))` where `MyHook` implements the `Hook` interface.
func (hooks LevelHooks) Add(hook Hook) {
	for _, level := range hook.Levels() {
		hooks[level] = append(hooks[level], hook)
	}
}

// Fire all the hooks for the passed level. Used by `entry.log` to fire
// appropriate hooks for a log entry.
func (hooks LevelHooks) Fire(level logger.Level, log []logger.Field) error {
	for _, hook := range hooks[level] {
		if err := hook.Fire(log); err != nil {
			return err
		}
	}

	return nil
}

type FileHook struct {
	Writer *os.File
}

// Creates a hook to be added to an instance of logger. This is called with
// `hook, err := NewFileHook("udp", "localhost:514", syslog.LOG_DEBUG, "")`
// `if err == nil { log.Hooks.Add(hook) }`
func NewFileHook(name string) (*FileHook, error) {
	file, err := os.Create(name)
	if err == nil {
		return &FileHook{Writer: file}, err
	}
	return nil, err
}
