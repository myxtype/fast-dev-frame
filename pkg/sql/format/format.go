package format

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// scan for scanner helper
func Scan(data interface{}, value interface{}) error {
	if value == nil {
		return nil
	}
	switch value.(type) {
	case []byte:
		return json.Unmarshal(value.([]byte), data)
	case string:
		return json.Unmarshal([]byte(value.(string)), data)
	default:
		return fmt.Errorf("val type is valid, is %+v", value)
	}
}

// for valuer helper
func Value(data interface{}) (interface{}, error) {
	vi := reflect.ValueOf(data)
	if vi.IsZero() {
		return nil, nil
	}
	return json.Marshal(data)
}
