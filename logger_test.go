package logger_test

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"strings"
	"testing"
	"time"

	loggerNoriCommon "github.com/nori-io/nori-common/logger"
	"github.com/stretchr/testify/assert"

	"github.com/nori-io/logger"
)

func TestLogger(t *testing.T) {
	a := assert.New(t)
	testData := "test"

	type decodedData struct {
		Level string    `json:"level"`
		Msg   string    `json:"msg"`
		Time  time.Time `json:"time"`
	}

	buferSize := 100
	buf := bytes.Buffer{}
	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))
	logTest1.Log(loggerNoriCommon.LevelInfo, testData)
	result := make([]byte, buferSize)
	_, err := buf.Read(result)

	decodedDataTest := new(decodedData)
	result = []byte(strings.TrimRight(string(result), "\x00"))
	err = json.Unmarshal(result, &decodedDataTest)
	testResult := &decodedData{
		Level: "info",
		Msg:   "test",
		Time:  decodedDataTest.Time,
	}
	a.Equal(testResult, decodedDataTest)
	a.NoError(err)
	buf.Reset()

	/*	testResult = "[FATAL]{\"Msg\":\"test\"}"
		logTest1.Fatal("%s", []byte(testData))
		_, err = buf.Read(result)
		a.Equal(testResult, string(result))
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
		buf.Reset()*/

}

func TestLoggerWith(t *testing.T) {
	a := assert.New(t)
	buferSize := 250
	result := make([]byte, buferSize)
	result2 := make([]byte, buferSize)
	buf := bytes.Buffer{}

	hook, _ := logger.NewFileHook("logger1")
	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf), logger.SetFileHook(*hook))
	logTest1.Log(loggerNoriCommon.LevelInfo, "test")
	buf.Read(result)
	buf.Reset()

	logTest2 := logTest1.With(loggerNoriCommon.Field{Key: "1", Value: "test1"}, loggerNoriCommon.Field{Key: "2", Value: "test2"})
	logTest2.Log(loggerNoriCommon.LevelInfo, "test")
	buf.Read(result2)
	a.Equal(false, &logTest1 == &logTest2)
	a.Equal(false, string(result) == string(result2))

	buf.Reset()
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
