package logger_test

import (
	"bytes"
	//"sync"
	"testing"

	//"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	loggerNoriCommon "github.com/nori-io/nori-common/logger"

	"github.com/nori-io/logger"
)

func TestLogger(t *testing.T) {
	a := assert.New(t)
	testData := "test"
	buferSize := 4
	buf := bytes.Buffer{}
	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))
	logTest1.Log(loggerNoriCommon.LevelInfo, testData)
	result := make([]byte, buferSize)
	_, err := buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	logTest1.Fatal("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	logTest1.Panic("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	logTest1.Error("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	logTest1.Critical("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	logTest1.Debug("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	logTest1.Info("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	logTest1.Notice("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	logTest1.Warning("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

}

func TestLoggerWith(t *testing.T) {
	a := assert.New(t)
	buferSize := 71
	result := make([]byte, buferSize)
	result2 := make([]byte, buferSize)
	buf := bytes.Buffer{}

	//logHook:= logger.New()

	//hook, err := syslog.NewSyslogHook("", "", syslog2.LOG_INFO, "")
	//	if err == nil {
	//logHook.
	//}

	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf), logger.SetHook(nil))
	logTest1.Log(loggerNoriCommon.LevelInfo, "test")
	buf.Read(result)
	buf.Reset()

	logTest2 := logTest1.With(loggerNoriCommon.Field{Key: "1", Value: "test1"}, loggerNoriCommon.Field{Key: "2", Value: "test2"})
	logTest2.Log(loggerNoriCommon.LevelInfo, "test")
	a.Equal(false, &logTest1 == &logTest2)
	a.Equal(false, string(result) == string(result2))
	buf.Read(result2)

	buf.Reset()
}
