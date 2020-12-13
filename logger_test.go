package logger_test

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	loggerNoriCommon "github.com/nori-io/common/v4/pkg/domain/logger"
	"github.com/nori-io/logger"
	"github.com/stretchr/testify/assert"
)

type decodedData struct {
	Level string    `json:"level"`
	Msg   string    `json:"msg"`
	Time  time.Time `json:"time"`
}

func TestLogger_Log(t *testing.T) {
	a := assert.New(t)

	var (
		data = &decodedData{
			Level: "error",
			Msg:   "test",
		}
		decoded = &decodedData{}
		buf     = bytes.Buffer{}
	)

	l := logger.New(logger.SetJsonFormatter(""), logger.SetOutWriter(&buf))
	l.Log(loggerNoriCommon.LevelInfo, data.Msg)

	result := []byte(strings.TrimRight(buf.String(), "\x00"))
	err := json.Unmarshal(result, &decoded)
	a.NoError(err)
	data.Time = decoded.Time

	a.Equal(data, decoded)
}

func TestLogger_Fatal(t *testing.T) {
	if os.Getenv("BE_FATAL") == "1" {
		l := logger.New()
		l.Fatal("%s", "fatal")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLogger_Fatal")
	cmd.Env = append(os.Environ(), "BE_FATAL=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestLogger_Panic(t *testing.T) {
	a := assert.New(t)

	var (
		data = &decodedData{
			Level: "panic",
			Msg:   "test",
		}
		decoded = &decodedData{}
		buf     = bytes.Buffer{}
	)

	l := logger.New(logger.SetJsonFormatter(""), logger.SetOutWriter(&buf))

	var recovered interface{}
	func() {
		defer func() {
			recovered = recover()
		}()
		l.Panic(data.Msg)
	}()
	a.NotNil(recovered)

	result := []byte(strings.TrimRight(buf.String(), "\x00"))
	err := json.Unmarshal(result, &decoded)
	a.NoError(err)

	data.Time = decoded.Time
	a.Equal(data, decoded)
}

func TestLogger_Error(t *testing.T) {
	a := assert.New(t)

	var (
		data = &decodedData{
			Level: "error",
			Msg:   "test",
		}
		decoded = &decodedData{}
		buf     = bytes.Buffer{}
	)

	l := logger.New(logger.SetJsonFormatter(""), logger.SetOutWriter(&buf))

	l.Error(data.Msg)

	record := []byte(strings.TrimRight(buf.String(), "\x00"))
	err := json.Unmarshal(record, decoded)
	a.NoError(err)

	data.Time = decoded.Time

	a.Equal(data, decoded)
}

func TestLogger_Critical(t *testing.T) {
	a := assert.New(t)
	var (
		data = &decodedData{
			Level: "critical",
			Msg:   "test",
		}
		decoded = &decodedData{}
		buf     = bytes.Buffer{}
	)

	l := logger.New(logger.SetJsonFormatter(""), logger.SetOutWriter(&buf))

	l.Critical(data.Msg)

	result := []byte(strings.TrimRight(buf.String(), "\x00"))
	err := json.Unmarshal(result, &decoded)
	a.NoError(err)

	data.Time = decoded.Time

	a.Equal(data, decoded)
}

func TestLogger_Debug(t *testing.T) {
	a := assert.New(t)

	var (
		data = &decodedData{
			Level: "debug",
			Msg:   "test",
		}
		decoded = &decodedData{}
		buf     = bytes.Buffer{}
	)

	l := logger.New(logger.SetJsonFormatter(""), logger.SetOutWriter(&buf))

	l.Debug(data.Msg)

	result := []byte(strings.TrimRight(buf.String(), "\x00"))
	err := json.Unmarshal(result, &decoded)
	a.NoError(err)

	data.Time = decoded.Time

	a.Equal(data, decoded)
}

func TestLogger_Info(t *testing.T) {
	a := assert.New(t)

	var (
		data = &decodedData{
			Level: "info",
			Msg:   "test",
		}
		decoded = &decodedData{}
		buf     = bytes.Buffer{}
	)

	l := logger.New(logger.SetJsonFormatter(""), logger.SetOutWriter(&buf))

	l.Info(data.Msg)

	result := []byte(strings.TrimRight(buf.String(), "\x00"))
	err := json.Unmarshal(result, &decoded)
	a.NoError(err)

	a.Equal(data, decoded)
}

func TestLogger_Notice(t *testing.T) {
	a := assert.New(t)

	var (
		data = &decodedData{
			Level: "notice",
			Msg:   "test",
		}
		decoded = &decodedData{}
		buf     = bytes.Buffer{}
	)

	l := logger.New(logger.SetJsonFormatter(""), logger.SetOutWriter(&buf))

	l.Notice(data.Msg)

	result := []byte(strings.TrimRight(buf.String(), "\x00"))
	err := json.Unmarshal(result, &decoded)
	a.NoError(err)

	a.Equal(data, decoded)
}

func TestLogger_Warning(t *testing.T) {
	a := assert.New(t)

	var (
		data = &decodedData{
			Level: "warning",
			Msg:   "test",
		}
		decoded = &decodedData{}
		buf     = bytes.Buffer{}
	)

	l := logger.New(logger.SetJsonFormatter(""), logger.SetOutWriter(&buf))

	l.Warning(data.Msg)

	result := []byte(strings.TrimRight(buf.String(), "\x00"))
	err := json.Unmarshal(result, &decoded)
	a.NoError(err)

	a.Equal(data, decoded)
}

func TestLoggerWith(t *testing.T) {
	a := assert.New(t)

	var (
		buf        = bytes.Buffer{}
		recA, recB string
	)

	l1 := logger.New(logger.SetJsonFormatter(""), logger.SetOutWriter(&buf))
	l1.Log(loggerNoriCommon.LevelInfo, "test")
	recA = buf.String()
	buf.Reset()

	l2 := l1.With(loggerNoriCommon.Field{Key: "foo", Value: "bar"}, loggerNoriCommon.Field{Key: "key", Value: "value"})
	l2.Log(loggerNoriCommon.LevelInfo, "test")
	recB = buf.String()

	a.False(&l1 == &l2)
	a.NotEqual(recA, recB)
}
