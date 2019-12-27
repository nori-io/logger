package logger

import (
	"bytes"
	"encoding/json"
	"fmt"

	log "github.com/nori-io/nori-common/logger"
)

type fieldKey string

type Formatter interface {
	Format(field ...log.Field) ([]byte, error)
}

type FieldMap map[fieldKey]string

func (f FieldMap) resolve(key fieldKey) string {
	if k, ok := f[key]; ok {
		return k
	}
	return string(key)
}

type JSONFormatter struct {
	DataKey  string
	FieldMap FieldMap
}

func (f *JSONFormatter) FormatFields(fields ...log.Field) ([]byte, error) {
	data := make(map[string]string, 1)
	for _, v := range fields {

		data[v.Key] = v.Value

	}
	var b *bytes.Buffer
	b = &bytes.Buffer{}
	encoder := json.NewEncoder(b)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return b.Bytes(), nil
}
