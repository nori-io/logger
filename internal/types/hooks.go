package types

import (
	"github.com/nori-io/common/v4/pkg/domain/logger"
)

type LevelHooks map[logger.Level][]logger.Hook

func (hooks LevelHooks) Add(hook logger.Hook) {
	for _, level := range hook.Levels() {
		hooks[level] = append(hooks[level], hook)
	}
}

func (hooks LevelHooks) Fire(entry logger.Entry, fields ...logger.Field) error {
	for _, hook := range hooks[entry.Level] {
		if err := hook.Fire(entry, fields...); err != nil {
			return err
		}
	}
	return nil
}
