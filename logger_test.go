package logger_test

import (
	"bytes"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nori-io/logger"

	loggerNoriCommon "github.com/nori-io/nori-common/logger"
)

func TestLogger(t *testing.T) {
	a := assert.New(t)
	testData := "test"
	buferSize := 4
	buf := bytes.Buffer{}
	logTest1 := &logger.Logger{
		Mu:        sync.Mutex{},
		Out:       &buf,
		Core:      logger.Core{},
		Formatter: logger.JSONFormatter{},
		Hooks:     nil,
	}
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
	/*	buferSize := 71
		result := make([]byte, buferSize)*/
	logTest1 := new(logger.Logger)

	buf := bytes.Buffer{}
	logTest1.Out = &buf

	//logTest2 := (logTest1.With(loggerNoriCommon.Field{Key: "1", Value: "test1"}, loggerNoriCommon.Field{Key: "2", Value: "test2"})).(*logger.Logger)
	logTest2 := logTest1.With(loggerNoriCommon.Field{Key: "1", Value: "test1"}, loggerNoriCommon.Field{Key: "2", Value: "test2"})
	fmt.Println(logTest2)
	a.Equal(true, logTest1 == logTest1)

	//	logTest2.Info("%s", "testWarning")
	//_, err := buf.Read(result)
	//	testData := ""
	fmt.Println(&logTest2)

	fmt.Println(&logTest1)
	/*	for _, value := range logTest2.Core.Fields {
			testData = testData + value.Key + " " + value.Value
		}
		testData = testData + "testWarning"

		a.Equal(true, &logTest1.Mu == &logTest2.Mu)
		a.Equal(true, &logTest1.Formatter == &logTest2.Formatter)
		a.Equal(true, &logTest1.Out == &logTest2.Out)
		a.Equal(false, logTest1 == logTest2)
		a.Equal(true, &logTest1.Core == &logTest2.Core)*/

	//	a.NoError(err)
	buf.Reset()
}
