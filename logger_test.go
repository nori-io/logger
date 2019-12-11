package logger_test

import (
	"bytes"
	"fmt"
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
	log := logger.Logger{
		Out:  &buf,
		Core: logger.Core{Fields: []loggerNoriCommon.Field{}},
	}

	log.Log(loggerNoriCommon.LevelInfo, testData)
	result := make([]byte, buferSize)
	_, err := buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	log.Printf("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	log.Fatal("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	log.Panic("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	log.Error("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	log.Critical("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	log.Debug("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	log.Info("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	log.Notice("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	log.Warning("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, string(result))
	a.NoError(err)
	buf.Reset()

	//testFields := []loggerNoriCommon.Field{loggerNoriCommon.Field{Key: "1", Value: "test1"}, loggerNoriCommon.Field{Key: "2", Value: "test2"}}
	buferSize = 25
	result = make([]byte, buferSize)

	log.With(loggerNoriCommon.Field{Key: "1", Value: "test1"}, loggerNoriCommon.Field{Key: "2", Value: "test2"})

	log.Warning("%s", "testWarning")

	_, err = buf.Read(result)
	fmt.Println("result is", string(result))
	fmt.Println("log.core.Fields", log.Core.Fields)

	str := ""

	for _, value := range log.Core.Fields {
		str = str + value.Key + " " + value.Value

	}
	str = str + "testWarning"
	a.Equal(string(result), str)
	a.NoError(err)
	buf.Reset()

}

func TestFormatter(t *testing.T) {
	//formatter := &logger.JSONFormatter{}

}
