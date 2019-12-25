package logger

import (
	"bytes"
	"encoding/json"
	"fmt"

	log "github.com/nori-io/nori-common/logger"
)

type fieldKey string

// FieldMap allows customization of the key names for default fields.
type FieldMap map[fieldKey]string

func (f FieldMap) resolve(key fieldKey) string {
	if k, ok := f[key]; ok {
		return k
	}

	return string(key)
}

// JSONFormatter formats logs into parsable json
type JSONFormatter struct {
	// TimestampFormat sets the format used for marshaling timestamps.
	TimestampFormat string

	// DisableTimestamp allows disabling automatic timestamps in output
	DisableTimestamp bool

	// DataKey allows users to put all the log entry parameters into a nested dictionary at a given key.
	DataKey  string
	FieldMap FieldMap

	//  PrettyPrint will indent all json logs
	PrettyPrint bool
}

func (f *JSONFormatter) FormatFields(fields ...log.Field) ([]byte, error) {

	data := make(map[string]string, 1)
	for _, v := range fields {

		data[v.Key] = v.Value

	}

	var b *bytes.Buffer
	b = &bytes.Buffer{}

	encoder := json.NewEncoder(b)
	if f.PrettyPrint {
		encoder.SetIndent("", "  ")
	}
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}

	return b.Bytes(), nil
}
