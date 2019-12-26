package logger_test

import (
	"bytes"
	"testing"

	loggerNoriCommon "github.com/nori-io/nori-common/logger"
	"github.com/stretchr/testify/assert"

	"github.com/nori-io/logger"
)

func TestLocalhostAddAndPrint(t *testing.T) {

	buf := bytes.Buffer{}
	a := assert.New(t)

	hook, err := logger.NewFileHook("test_file")
	if err != nil {
		t.Errorf("Can't create hook")
	}

	a.NoError(err)
	hook2, err := logger.NewFileHook("test_file2")
	if err != nil {
		t.Errorf("Can't create hook")
	}
	a.NoError(err)

	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))
	logTest1.AddHook(hook)
	hook2.Levels()
	logTest1.AddHook(hook2)
	logTest1.Info("testInfo")
	logTest2 := logTest1.With(loggerNoriCommon.Field{Key: "1", Value: "test1"}, loggerNoriCommon.Field{Key: "2", Value: "test2"})
	logTest2.Log(loggerNoriCommon.LevelInfo, "test")

	logTest2.Warning("done")

}
