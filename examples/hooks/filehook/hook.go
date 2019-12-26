package logger

import (
	"io/ioutil"
	"os"

	"github.com/nori-io/nori-common/logger"
)

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
func NewFileHookTest(path string, name string) (*FileHook, error) {
	file, err := ioutil.TempFile(path, name)
	if err == nil {
		return &FileHook{Writer: file}, err
	}
	return nil, err
}
