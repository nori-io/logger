package formatter_test

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	loggerNoriCommon "github.com/nori-io/nori-common/logger"

	logger "github.com/nori-io/logger/formatter/json"
)

func TestErrorNotLost(t *testing.T) {

	logTest1 := &logger.Logger{
		Out:       nil,
		Mu:        &sync.Mutex{},
		Fields:    make([]loggerNoriCommon.Field, 2),
		Formatter: &JSONFormatter{},
		Hooks:     nil,
	}

	testField := loggerNoriCommon.Field{Key: "key1", Value: "value1"}

	b, err := logTest1.Formatter.FormatFields(testField)
	b2, _ := logTest1.Formatter.FormatFields(loggerNoriCommon.Field{
		Key:   "msg",
		Value: fmt.Sprintf("test"),
	})

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

}
