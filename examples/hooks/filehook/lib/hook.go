package lib

import (
	"bytes"
	"os"

	"github.com/nori-io/common/v4/pkg/domain/logger"
)

type FileHook struct {
	Writer *os.File
}

func NewFileHook(name string) (logger.Hook, error) {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		return &FileHook{Writer: file}, err
	}
	return nil, err
}

func (hook *FileHook) Levels() []logger.Level {
	return []logger.Level{logger.LevelFatal, logger.LevelPanic, logger.LevelNotice, logger.LevelCritical, logger.LevelError,
		logger.LevelWarning, logger.LevelInfo}
}

func (hook *FileHook) Fire(entry logger.Entry, field ...logger.Field) error {
	if entry.Level == logger.LevelDebug {
		return nil
	}

	b := bytes.Buffer{}
	out, _ := entry.Formatter.Format(entry, field...)
	b.Write(out)
	_, err := hook.Writer.Write(b.Bytes())
	return err
}
