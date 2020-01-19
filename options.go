package logger

import (
	"io"

	"github.com/nori-io/logger/pkg/formatters"
)

type Option interface {
	apply(*Logger)
}

type optionFunc func(*Logger)

func (f optionFunc) apply(log *Logger) {
	f(log)
}

func SetJsonFormatter() Option {
	return optionFunc(func(log *Logger) {
		log.Formatter = &formatters.JSONFormatter{}
	})
}

func SetOutWriter(writer io.Writer) Option {
	return optionFunc(func(log *Logger) {
		log.Out = writer
	})
}
