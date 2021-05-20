package formatters_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/nori-io/logger/pkg/formatters"
	"github.com/stretchr/testify/assert"

	"github.com/nori-io/common/v5/pkg/domain/logger"
)

func TestJSONFormatter_Format(t *testing.T) {
	a := assert.New(t)

	type dataOne struct {
		Level     string `json:"level"`
		Msg       string `json:"msg"`
		Ts        string `json:"ts"`
		Component string `json:"component"`
	}

	level := logger.LevelInfo
	msg := "foo"
	ts := time.Now()
	key := "component"
	value := "test"

	exp1 := dataOne{
		Level:     level.String(),
		Msg:       msg,
		Ts:        ts.Format(time.RFC3339),
		Component: value,
	}

	f := formatters.JSONFormatter{
		TimeFormat: time.RFC3339,
	}

	b, err := f.Format(logger.Entry{
		Level:   level,
		Time:    ts,
		Message: msg,
	}, logger.Field{
		Key:   key,
		Value: value,
	})

	if err != nil {
		t.Fatal("formatter.Format returned error: " + err.Error())
	}

	var src1 = new(dataOne)
	err = json.Unmarshal(b, src1)
	a.NoError(err, "Unable to unmarshal formatted entry")

	a.EqualValues(exp1, *src1)
}
