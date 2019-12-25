package logger_test

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	loggerNoriCommon "github.com/nori-io/nori-common/logger"

	"github.com/nori-io/logger"
)

func TestErrorNotLost(t *testing.T) {

	logTest1 := &logger.Logger{
		Out:       nil,
		Mu:        &sync.Mutex{},
		Fields:    make([]loggerNoriCommon.Field, 2),
		Formatter: &logger.JSONFormatter{},
		Hooks:     nil,
	}

	testField := loggerNoriCommon.Field{Key: "key1", Value: "value1"}

	b, err := logTest1.Formatter.FormatFields(testField)
	b2, _ := logTest1.Formatter.FormatMessage(loggerNoriCommon.Field{
		Key:   "msg",
		Value: fmt.Sprintf("test"),
	})
	fmt.Println("b", string(b))
	fmt.Println("b2", string(b2))
	type decodedData struct {
		Key  string    `json:"key1"`
		Time time.Time `json:"time"`
	}
	type decodedData2 struct {
		Msg  string    `json:"msg"`
		Time time.Time `json:"time"`
	}

	decodedDataTest := new(decodedData)
	decodedDataTest2 := new(decodedData2)

	err = json.Unmarshal(b, &decodedDataTest)
	if err != nil {
		t.Fatal("Unable to unmarshal formatted entry: ", err)
	}

	err = json.Unmarshal(b2, &decodedDataTest2)
	if err != nil {
		t.Fatal("Unable to unmarshal formatted entry: ", err)
	}
	fmt.Println("decodedDataTest", decodedDataTest)

	fmt.Println("decodedDataTest", decodedDataTest2)

}
