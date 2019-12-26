package logger_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	loggerNoriCommon "github.com/nori-io/nori-common/logger"
	"github.com/stretchr/testify/assert"

	"github.com/nori-io/logger"
)

func TestLocalhostAddAndPrint(t *testing.T) {

	buf := bytes.Buffer{}
	a := assert.New(t)

	hook, err := NewFileHookTest("","file_test")
	if err != nil {
		t.Errorf("Can't create hook")
	}

	a.NoError(err)
	hook2, err := NewFileHookTest("","file_test2")
	if err != nil {
		t.Errorf("Can't create hook")
	}
	a.NoError(err)

	logTest1 := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))
	logTest1.AddHook(hook)
	hook2.Levels()
	logTest1.AddHook(hook2)
	logTest1.Info("testInfo")
	logTest2 := logTest1.With(loggerNoriCommon.Field{Key: "1", Value: "test1"}, loggerNoriCommon.Field{Key: "2", Value: "test2"})
	logTest2.Log(loggerNoriCommon.LevelInfo, "test")

	logTest2.Warning("done")

	fileTest1, err1 := os.Open("file_test")
	if err1 != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fileTest1.Close()

	testData := "{\"level\":\"info\",\"msg\":\"testInfo\"}\n"
	testData2 := "{\"1\":\"test1\",\"2\":\"test2\",\"level\":\"info\",\"msg\":\"test\"}\n"
	testData3 := "{\"1\":\"test1\",\"2\":\"test2\",\"level\":\"warning\",\"msg\":\"done\"}\n"

	rows := make([]string, 3)
	r := bufio.NewReader(fileTest1)
	for i := 0; i < 3; i++ {
		rows[i], err = r.ReadString(10) //0x0A separator = newline
		if err == io.EOF {
			//do something here
			break
		}
	}
	a.Equal(rows[0], testData)
	a.Equal(rows[1], testData2)
	a.Equal(rows[2], testData3)
	fileTest2, err2 := os.Open("file_test2")
	if err2 != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fileTest2.Close()

	r = bufio.NewReader(fileTest2)
	for i := 0; i < 3; i++ {
		rows[i], err = r.ReadString(10) //0x0A separator = newline
		if err == io.EOF {
			//do something here
			break
		}
	}
	a.Equal(rows[0], testData)
	a.Equal(rows[1], testData2)
	a.Equal(rows[2], testData3)

}

func NewFileHookTest(path string, name string) (*logger.FileHook, error) {
	file, err := ioutil.TempFile(path,name)
	if err == nil {
		return &logger.FileHook{Writer: file}, err
	}
	return nil, err
}