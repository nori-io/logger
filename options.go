package logger

import (
	"github.com/nori-io/logger/formatter"
	"io"

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
		log.Formatter = &formatter.JSONFormatter{}
	})
}

func SetOutWriter(writer io.Writer) Option {
	return optionFunc(func(log *Logger) {
		log.Out = writer
	})
}