package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/nori-io/nori-common/logger"
)

type JSONFormatter struct {}

func (f *JSONFormatter) Format(msg string, ts string, fields ...log.Field) ([]byte, error) {
	data := make(map[string]string, len(fields) + 2)
	for _, v := range fields {
		data[v.Key] = v.Value
	}
	data["msg"] = msg
	data["ts"] = ts

	var b *bytes.Buffer
	b = &bytes.Buffer{}
	encoder := json.NewEncoder(b)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return b.Bytes(), nil
}

