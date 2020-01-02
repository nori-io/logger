package logger_test

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	loggerNoriCommon "github.com/nori-io/nori-common/logger"
	"github.com/stretchr/testify/assert"

	"github.com/nori-io/logger"
)

func TestLogger_Log(t *testing.T) {
	a := assert.New(t)
	testData := "test"

	type decodedData struct {
		Level string    `json:"level"`
		Msg   string    `json:"msg"`
		Time  time.Time `json:"time"`
	}

	bufferSize := 100
	buf := bytes.Buffer{}
	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))
	logTest1.Log(loggerNoriCommon.LevelInfo, testData)
	result := make([]byte, bufferSize)
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
}

func TestLogger_Fatal(t *testing.T) {
	a := assert.New(t)
	testData := "test"
	bufferSize := 110
	buf := bytes.Buffer{}

	type decodedData struct {
		Level string    `json:"level"`
		Msg   string    `json:"msg"`
		Time  time.Time `json:"time"`
	}
	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))

	logTest1.Fatal("%s", testData)
	result := make([]byte, bufferSize)
	_, err := buf.Read(result)

	decodedDataTest := new(decodedData)
	result = []byte(strings.TrimRight(string(result), "\x00"))
	err = json.Unmarshal(result, &decodedDataTest)
	testResult := &decodedData{
		Level: "fatal",
		Msg:   "test",
		Time:  decodedDataTest.Time,
	}
	a.Equal(testResult, decodedDataTest)
	a.NoError(err)
	buf.Reset()
}

func TestLogger_Panic(t *testing.T) {
	a := assert.New(t)
	testData := "test"
	bufferSize := 110
	buf := bytes.Buffer{}

	type decodedData struct {
		Level string    `json:"level"`
		Msg   string    `json:"msg"`
		Time  time.Time `json:"time"`
	}
	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))

	logTest1.Panic(testData)
	result := make([]byte, bufferSize)
	_, err := buf.Read(result)

	decodedDataTest := new(decodedData)
	result = []byte(strings.TrimRight(string(result), "\x00"))
	err = json.Unmarshal(result, &decodedDataTest)
	testResult := &decodedData{
		Level: "panic",
		Msg:   "test",
		Time:  decodedDataTest.Time,
	}
	a.Equal(testResult, decodedDataTest)
	a.NoError(err)
	buf.Reset()
}

func TestLogger_Error(t *testing.T) {
	a := assert.New(t)
	testData := "test"
	bufferSize := 110
	buf := bytes.Buffer{}

	type decodedData struct {
		Level string    `json:"level"`
		Msg   string    `json:"msg"`
		Time  time.Time `json:"time"`
	}
	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))

	logTest1.Error(testData)
	result := make([]byte, bufferSize)
	_, err := buf.Read(result)

	decodedDataTest := new(decodedData)
	result = []byte(strings.TrimRight(string(result), "\x00"))
	err = json.Unmarshal(result, &decodedDataTest)
	testResult := &decodedData{
		Level: "error",
		Msg:   "test",
		Time:  decodedDataTest.Time,
	}
	a.Equal(testResult, decodedDataTest)
	a.NoError(err)
	buf.Reset()
}

func TestLogger_Critical(t *testing.T) {
	a := assert.New(t)
	testData := "test"
	bufferSize := 110
	buf := bytes.Buffer{}

	type decodedData struct {
		Level string    `json:"level"`
		Msg   string    `json:"msg"`
		Time  time.Time `json:"time"`
	}
	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))

	logTest1.Critical(testData)
	result := make([]byte, bufferSize)
	_, err := buf.Read(result)

	decodedDataTest := new(decodedData)
	result = []byte(strings.TrimRight(string(result), "\x00"))
	err = json.Unmarshal(result, &decodedDataTest)
	testResult := &decodedData{
		Level: "critical",
		Msg:   "test",
		Time:  decodedDataTest.Time,
	}
	a.Equal(testResult, decodedDataTest)
	a.NoError(err)
	buf.Reset()

}

func TestLogger_Debug(t *testing.T) {
	a := assert.New(t)
	testData := "test"
	bufferSize := 110
	buf := bytes.Buffer{}

	type decodedData struct {
		Level string    `json:"level"`
		Msg   string    `json:"msg"`
		Time  time.Time `json:"time"`
	}
	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))

	logTest1.Debug(testData)
	result := make([]byte, bufferSize)
	_, err := buf.Read(result)

	decodedDataTest := new(decodedData)
	result = []byte(strings.TrimRight(string(result), "\x00"))
	err = json.Unmarshal(result, &decodedDataTest)
	testResult := &decodedData{
		Level: "debug",
		Msg:   "test",
		Time:  decodedDataTest.Time,
	}
	a.Equal(testResult, decodedDataTest)
	a.NoError(err)
	buf.Reset()
}

func TestLogger_Info(t *testing.T) {
	a := assert.New(t)
	testData := "test"
	bufferSize := 110
	buf := bytes.Buffer{}

	type decodedData struct {
		Level string    `json:"level"`
		Msg   string    `json:"msg"`
		Time  time.Time `json:"time"`
	}
	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))

	logTest1.Info(testData)
	result := make([]byte, bufferSize)
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
}

func TestLogger_Notice(t *testing.T) {
	a := assert.New(t)
	testData := "test"
	bufferSize := 110
	buf := bytes.Buffer{}

	type decodedData struct {
		Level string    `json:"level"`
		Msg   string    `json:"msg"`
		Time  time.Time `json:"time"`
	}
	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))

	logTest1.Notice(testData)
	result := make([]byte, bufferSize)
	_, err := buf.Read(result)

	decodedDataTest := new(decodedData)
	result = []byte(strings.TrimRight(string(result), "\x00"))
	err = json.Unmarshal(result, &decodedDataTest)
	testResult := &decodedData{
		Level: "notice",
		Msg:   "test",
		Time:  decodedDataTest.Time,
	}
	a.Equal(testResult, decodedDataTest)
	a.NoError(err)
	buf.Reset()
}

func TestLogger_Warning(t *testing.T) {
	a := assert.New(t)
	testData := "test"
	bufferSize := 110
	buf := bytes.Buffer{}

	type decodedData struct {
		Level string    `json:"level"`
		Msg   string    `json:"msg"`
		Time  time.Time `json:"time"`
	}
	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))

	logTest1.Warning(testData)
	result := make([]byte, bufferSize)
	_, err := buf.Read(result)

	decodedDataTest := new(decodedData)
	result = []byte(strings.TrimRight(string(result), "\x00"))
	err = json.Unmarshal(result, &decodedDataTest)
	testResult := &decodedData{
		Level: "warning",
		Msg:   "test",
		Time:  decodedDataTest.Time,
	}
	a.Equal(testResult, decodedDataTest)
	a.NoError(err)
	buf.Reset()

}

func TestLoggerWith(t *testing.T) {
	a := assert.New(t)
	buferSize := 250
	result := make([]byte, buferSize)
	result2 := make([]byte, buferSize)
	buf := bytes.Buffer{}

	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))
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
