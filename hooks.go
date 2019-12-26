package logger

import (
	"github.com/nori-io/nori-common/logger"
)

type LevelHooks map[logger.Level][]logger.Hook

func (hooks LevelHooks) Add(hook logger.Hook) {
	for _, level := range hook.Levels() {
		hooks[level] = append(hooks[level], hook)
	}
}

func (hooks LevelHooks) Fire(level logger.Level, message []byte) error {
	for _, hook := range hooks[level] {
		if err := hook.Fire(level, message); err != nil {
			return err
		}
	}
	return nil
}
