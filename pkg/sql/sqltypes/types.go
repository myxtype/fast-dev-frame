package sqltypes

import (
	"database/sql/driver"
	"frame/pkg/sql/format"
)

// Int64数组
type Int64Array []int64

func (t Int64Array) GormDataType() string {
	return "json"
}

func (t *Int64Array) Scan(src interface{}) error {
	return format.Scan(t, src)
}

func (t Int64Array) Value() (driver.Value, error) {
	return format.Value(t)
}

// 字符串数组
type StringArray []string

func (t StringArray) GormDataType() string {
	return "json"
}

func (t *StringArray) Scan(src interface{}) error {
	return format.Scan(t, src)
}

func (t StringArray) Value() (driver.Value, error) {
	return format.Value(t)
}

// 字典类型
type Map map[string]interface{}

func (t Map) GormDataType() string {
	return "json"
}

func (t *Map) Scan(src interface{}) error {
	return format.Scan(t, src)
}

func (t Map) Value() (driver.Value, error) {
	return format.Value(t)
}

// 字典数组
type List []Map

func (t List) GormDataType() string {
	return "json"
}

func (t *List) Scan(src interface{}) error {
	return format.Scan(t, src)
}

func (t List) Value() (driver.Value, error) {
	return format.Value(t)
}
