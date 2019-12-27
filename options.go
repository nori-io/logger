package logger

import (
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
		log.Formatter = &JSONFormatter{
			DataKey:  "",
			FieldMap: nil,
		}
	})
}

func SetOutWriter(writer io.Writer) Option {
	return optionFunc(func(log *Logger) {
		log.Out = &writer
	})
}