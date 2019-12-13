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
	DataKey string

	// FieldMap allows users to customize the names of keys for default fields.
	// As an example:
	// formatter := &JSONFormatter{
	//   	FieldMap: FieldMap{
	// 		 FieldKeyTime:  "@timestamp",
	// 		 FieldKeyLevel: "@level",
	// 		 FieldKeyMsg:   "@message",
	// 		 FieldKeyFunc:  "@caller",
	//    },
	// }
	FieldMap FieldMap

	// PrettyPrint will indent all json logs
	PrettyPrint bool
}

// Format renders a single log entry
func (f *JSONFormatter) Format(fields ...log.Field) ([]byte, error) {

	data := make([]log.Field, 1)
	for k, v := range fields {

		data[k] = v

	}

	if f.DataKey != "" {
		newData := make([]log.Field, 2)
		//newData[f.DataKey] = data
		data = newData
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	/*	if entry.err != "" {
			data[f.FieldMap.resolve(FieldKeyLogrusError)] = entry.err
		}
		if !f.DisableTimestamp {
			data[f.FieldMap.resolve(FieldKeyTime)] = entry.Time.Format(timestampFormat)
		}
		//data[f.FieldMap.resolve(FieldKeyMsg)] = log.Message
		data[f.FieldMap.resolve(FieldKeyLevel)] = logger.Level.String
		if entry.HasCaller() {
			data[f.FieldMap.resolve(FieldKeyFunc)] = entry.Caller.Function
			data[f.FieldMap.resolve(FieldKeyFile)] = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		}*/

	var b *bytes.Buffer
	/*if &[]log.Field!=nil {
		b = log.
	} else {*/
	b = &bytes.Buffer{}
	//}

	encoder := json.NewEncoder(b)
	if f.PrettyPrint {
		encoder.SetIndent("", "  ")
	}
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}

	return b.Bytes(), nil
}
