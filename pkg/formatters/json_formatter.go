package formatters

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/nori-io/common/v3/logger"
)

type JSONFormatter struct {
	TimeFormat string
}

func (f *JSONFormatter) Format(entry logger.Entry, fields ...logger.Field) ([]byte, error) {
	data := make(map[string]interface{}, len(fields)+3)
	for _, v := range fields {
		data[v.Key] = v.Value
	}
	data["level"] = entry.Level.String()
	data["ts"] = entry.Time.Format(f.TimeFormat)
	data["msg"] = entry.Message

	var b *bytes.Buffer
	b = &bytes.Buffer{}
	encoder := json.NewEncoder(b)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return b.Bytes(), nil
}
