package logger_test

import (
	"encoding/json"
	"sync"
	"testing"

	loggerNoriCommon "github.com/nori-io/nori-common/logger"

	"github.com/nori-io/logger"
)

func TestErrorNotLost(t *testing.T) {

	logTest1 := &logger.Logger{
		Out:  nil,
		Mu:   &sync.Mutex{},
		Core: logger.Core{},
		//Formatter: logger.JSONFormatter{},
		Hooks: nil,
	}

	testField := loggerNoriCommon.Field{Key: "key1", Value: "value1"}

	b, err := logTest1.Formatter.Format(testField)

	type data []loggerNoriCommon.Field

	//var jsonBlob data

	if err != nil {
		t.Fatal("Unable to format entry: ", err)
	}

	err = json.Unmarshal(b, &data{})
	if err != nil {
		t.Fatal("Unable to unmarshal formatted entry: ", err)
	}
	/*
		if entry["error"] != "wild walrus" {
			t.Fatal("Error field not set")
		}*/
}
