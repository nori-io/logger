package logger

import (
	"os"

	"github.com/nori-io/nori-common/logger"
)

type FileHook struct {
	Writer *os.File
}

func NewFileHook(name string) (logger.Hook, error) {
	file, err := os.Create(name)
	if err == nil {
		return &FileHook{Writer: file}, err
	}
	return nil, err
}

func (hook *FileHook) Levels() []logger.Level {
	return []logger.Level{logger.LevelFatal, logger.LevelPanic, logger.LevelNotice, logger.LevelCritical, logger.LevelError,
		logger.LevelWarning, logger.LevelInfo}
}

func (hook *FileHook) Fire(level logger.Level, message []byte) error {
	if level == logger.LevelDebug {
		return nil
	}
	_, err := hook.Writer.Write(message)
	return err
}
