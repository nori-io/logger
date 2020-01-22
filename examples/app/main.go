package main

import (
	"os"

	"github.com/nori-io/logger"
	logger2 "github.com/nori-io/nori-common/v2/logger"
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
