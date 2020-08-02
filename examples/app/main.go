package main

import (
	"os"

	logger2 "github.com/nori-io/common/v3/logger"
	"github.com/nori-io/logger"
)

func main() {
	const (
		component = "main.app"
	)
	var (
		field = logger2.Field{
			Key:   "component",
			Value: "examples.app",
		}
	)
	l := logger.New(logger.SetOutWriter(os.Stdout))
	l.Info("Info message from %s", component)

	l2 := l.With(field)
	l2.Info("Info message from %s with fields", component)
}
