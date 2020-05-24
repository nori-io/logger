package lib_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"testing"
	"time"

	loggerNoriCommon "github.com/nori-io/nori-common/v2/logger"
	"github.com/stretchr/testify/assert"

	"github.com/nori-io/logger"
)

func TestFileHook(t *testing.T) {
	a := assert.New(t)

	tmpFile, err := ioutil.TempFile("", "file_hook")
	defer os.Remove(tmpFile.Name()) // clean up

	hook, err := NewFileHook(tmpFile.Name())
	a.NoError(err, "Can't create hook")

	var (
		msg     = "foo bar"
		msg2    = "lorem ipsum"
		warning = "warning"
		key1    = "one"
		key2    = "two"
		val1    = "1"
		val2    = "2"
	)

	buf := bytes.Buffer{}
	log1 := logger.New(logger.SetJsonFormatter(""), logger.SetOutWriter(&buf))
	log1.AddHook(hook)
	log1.Info(msg)
	log2 := log1.With(loggerNoriCommon.Field{Key: key1, Value: val1}, loggerNoriCommon.Field{Key: key2, Value: val2})
	log2.Log(loggerNoriCommon.LevelInfo, msg2)
	log2.Warning(warning)

	file1, err1 := os.Open(tmpFile.Name())
	if err1 != nil {
		os.Exit(1)
	}

	type decodedData struct {
		Level string    `json:"level"`
		Msg   string    `json:"msg"`
		Time  time.Time `json:"ts"`
	}
	type decodedData2 struct {
		One   string    `json:"one"`
		Two   string    `json:"two"`
		Level string    `json:"level"`
		Msg   string    `json:"msg"`
		Time  time.Time `json:"ts"`
	}
	decodedDataTest := new(decodedData)
	decodedDataTest2 := new(decodedData2)

	testData1 := decodedData{
		Level: loggerNoriCommon.LevelInfo.String(),
		Msg:   msg,
		Time:  time.Time{},
	}
	testData2 := decodedData2{
		One:   val1,
		Two:   val2,
		Level: loggerNoriCommon.LevelInfo.String(),
		Msg:   msg2,
		Time:  time.Time{},
	}
	testData3 := decodedData2{
		One:   val1,
		Two:   val2,
		Level: loggerNoriCommon.LevelWarning.String(),
		Msg:   warning,
		Time:  time.Time{},
	}

	rows := make([]string, 3)
	r := bufio.NewReader(file1)
	for i := 0; i < 3; i++ {
		rows[i], err = r.ReadString('\n')
		if err == io.EOF {
			break
		}
	}
	err = json.Unmarshal([]byte(rows[0]), &decodedDataTest)
	a.NoError(err)
	a.Equal(testData1.Level, decodedDataTest.Level)
	a.Equal(testData1.Msg, decodedDataTest.Msg)

	err = json.Unmarshal([]byte(rows[1]), &decodedDataTest2)
	a.NoError(err)
	a.Equal(testData2.Level, decodedDataTest2.Level)
	a.Equal(testData2.Msg, decodedDataTest2.Msg)
	a.Equal(testData2.One, decodedDataTest2.One)
	a.Equal(testData2.Two, decodedDataTest2.Two)

	err = json.Unmarshal([]byte(rows[2]), &decodedDataTest2)
	a.NoError(err)
	a.Equal(testData3.Level, decodedDataTest2.Level)
	a.Equal(testData3.Msg, decodedDataTest2.Msg)
	a.Equal(testData3.One, decodedDataTest2.One)
	a.Equal(testData3.Two, decodedDataTest2.Two)

}
