package logger_test

import (
	"bytes"
	"testing"

	loggerNoriCommon "github.com/nori-io/nori-common/logger"
	"github.com/stretchr/testify/assert"

	"github.com/nori-io/logger"
)

func TestNew(t *testing.T) {
	// тут у тебя создается твой логгер
	a := assert.New(t)
	bufferSize := 54

	// создается твой хук
	//hook := *new(logger.Hook)
	fields := make([]loggerNoriCommon.Field, 1)
	fields = append(fields, loggerNoriCommon.Field{Key: "1", Value: "test1"})
	levelHooks := logger.LevelHooks{}
	testHook := &TestHook{Fired: true}
	levelHooks.Add(testHook)
	levelHooks.Fire(loggerNoriCommon.LevelInfo, fields)

	result := make([]byte, bufferSize)
	result2 := make([]byte, bufferSize)
	buf := bytes.Buffer{}

	//hook, err := syslog.NewSyslogHook("", "", syslog2.LOG_INFO, "")
	//	if err == nil {
	//logHook.
	//}

	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf), logger.SetHook(nil))
	logTest1.Info("test")
	//_, err:= buf.Read(result)
	//logTest1 := logger.New()
	//logTest1.Info("test")
	//	_, err = buf.Read(result2)
	a.Equal(string(result), string(result2))

}

type TestHook struct {
	Fired bool
}

func (hook *TestHook) Fire(fields []loggerNoriCommon.Field) error {
	hook.Fired = true
	return nil
}

func (hook *TestHook) Levels() []loggerNoriCommon.Level {
	return []loggerNoriCommon.Level{
		loggerNoriCommon.LevelFatal,
		loggerNoriCommon.LevelPanic,
		loggerNoriCommon.LevelNotice,
		loggerNoriCommon.LevelCritical,
		loggerNoriCommon.LevelError,
		loggerNoriCommon.LevelWarning,
		loggerNoriCommon.LevelInfo,
		loggerNoriCommon.LevelDebug,
	}
}
