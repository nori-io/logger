package logger

import (
	"io"
	"time"

	"github.com/nori-io/logger/pkg/formatters"
)

type Option interface {
	apply(*Logger)
}

type optionFunc func(*Logger)

func (f optionFunc) apply(log *Logger) {
	f(log)
}

func SetJsonFormatter(timeFormat string) Option {
	if timeFormat == "" {
		timeFormat = time.RFC3339
	}
	return optionFunc(func(log *Logger) {
		log.Formatter = &formatters.JSONFormatter{
			TimeFormat: timeFormat,
		}
	})
}

func SetOutWriter(writer io.Writer) Option {
	return optionFunc(func(log *Logger) {
		log.Out = writer
	})
}
