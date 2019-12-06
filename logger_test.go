package logger_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nori-io/logger"
)

func TestLogger(t *testing.T) {
	a := assert.New(t)
	testData := []byte("test")
	buferSize := 4
	buf := bytes.Buffer{}
	log := logger.Logger{
		Out:  &buf,
		Core: &logger.IoCore{},
	}

	log.Write([]byte(testData))
	result := make([]byte, buferSize)
	_, err := buf.Read(result)
	a.Equal(testData, result)
	a.NoError(err)
	buf.Reset()

	log.Printf("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, result)
	a.NoError(err)
	buf.Reset()

	log.Fatal("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, result)
	a.NoError(err)
	buf.Reset()

	log.Panic("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, result)
	a.NoError(err)
	buf.Reset()

	log.Error("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, result)
	a.NoError(err)
	buf.Reset()

	log.Critical("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, result)
	a.NoError(err)
	buf.Reset()

	log.Debug("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, result)
	a.NoError(err)
	buf.Reset()

	log.Info("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, result)
	a.NoError(err)
	buf.Reset()

	log.Notice("%s", []byte(testData))
	_, err = buf.Read(result)
	a.Equal(testData, result)
	a.NoError(err)
	buf.Reset()

	log.Warning("%s", []byte(testData))
	_, err = buf.Read(result)
	fmt.Println(result)
	a.Equal(testData, result)
	a.NoError(err)
	buf.Reset()

	var m map[string]interface{}
	m = make(map[string]interface{})
	m["test"] = "1"
	m["test2"] = "2"

	log.With(m)
	fmt.Println("log.Out", log.Out)

	fmt.Println(result)
	_, err = buf.Read(result)
	a.Equal(testData, result)
	a.NoError(err)
	buf.Reset()

}
