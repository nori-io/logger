package logger_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"
	"time"

	loggerNoriCommon "github.com/nori-io/nori-common/logger"
	"github.com/stretchr/testify/assert"

	"github.com/nori-io/logger"
	logger2 "github.com/nori-io/logger/examples/hooks/filehook"
)

func TestFileHook(t *testing.T) {

	buf := bytes.Buffer{}
	a := assert.New(t)

	hook, err := logger2.NewFileHook("file_test")
	if err != nil {
		t.Errorf("Can't create hook")
	}

	a.NoError(err)

	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))
	logTest1.AddHook(hook)
	logTest1.Info("testInfo")
	logTest2 := logTest1.With(loggerNoriCommon.Field{Key: "1", Value: "test1"}, loggerNoriCommon.Field{Key: "2", Value: "test2"})
	logTest2.Log(loggerNoriCommon.LevelInfo, "test")

	logTest2.Warning("done")

	fileTest1, err1 := os.Open("file_test")
	if err1 != nil {
		os.Exit(1)
	}
	defer os.Remove(fileTest1.Name())

	type decodedData struct {
		Level string    `json:"level"`
		Msg   string    `json:"msg"`
		Time  time.Time `json:"time"`
	}
	type decodedData2 struct {
		Num1  string    `json:"1"`
		Num2  string    `json:"2"`
		Level string    `json:"level"`
		Msg   string    `json:"msg"`
		Time  time.Time `json:"time"`
	}

	decodedDataTest := new(decodedData)
	decodedDataTest2 := new(decodedData2)

	testData1 := decodedData{
		Level: "info",
		Msg:   "testInfo",
		Time:  time.Time{},
	}
	testData2 := decodedData2{
		Num1:  "test1",
		Num2:  "test2",
		Level: "info",
		Msg:   "test",
		Time:  time.Time{},
	}

	testData3 := decodedData2{
		Num1:  "test1",
		Num2:  "test2",
		Level: "warning",
		Msg:   "done",
		Time:  time.Time{},
	}

	rows := make([]string, 3)
	r := bufio.NewReader(fileTest1)
	for i := 0; i < 3; i++ {
		rows[i], err = r.ReadString(10)
		if err == io.EOF {
			break
		}
	}

	err = json.Unmarshal([]byte(rows[0]), &decodedDataTest)
	a.NoError(err)
	a.Equal(decodedDataTest.Level, testData1.Level)
	a.Equal(decodedDataTest.Msg, testData1.Msg)

	err = json.Unmarshal([]byte(rows[1]), &decodedDataTest2)
	a.NoError(err)

	a.Equal(decodedDataTest2.Level, testData2.Level)
	a.Equal(decodedDataTest2.Msg, testData2.Msg)
	a.Equal(decodedDataTest2.Num1, testData2.Num1)
	a.Equal(decodedDataTest2.Num2, testData2.Num2)

	err = json.Unmarshal([]byte(rows[2]), &decodedDataTest2)
	a.NoError(err)

	a.Equal(decodedDataTest2.Level, testData3.Level)
	a.Equal(decodedDataTest2.Msg, testData3.Msg)
	a.Equal(decodedDataTest2.Num1, testData3.Num1)
	a.Equal(decodedDataTest2.Num2, testData3.Num2)

}
