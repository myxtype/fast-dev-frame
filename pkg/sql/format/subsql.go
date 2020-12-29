package format

import (
	"fmt"
	"strings"
)

// 子查询构造器
// 用户将多个子查询构造为一个查询
type SubSql struct {
	fields []string
	args   []interface{}
}

func NewSubSql() *SubSql {
	return &SubSql{}
}

// Add 添加一个子查询，fieldName表示子查询值的别名
func (s *SubSql) Add(fieldName, query string, args ...interface{}) {
	s.fields = append(s.fields, fmt.Sprintf("(%s) as `%s`", query, fieldName))
	if len(args) > 0 {
		s.args = append(s.args, args...)
	}
}

// Format 获取完整的查询语句
func (s *SubSql) Format() string {
	return fmt.Sprintf("SELECT %s", strings.Join(s.fields, ","))
}

// Args 获取查询的全部参数
func (s *SubSql) Args() []interface{} {
	return s.args
}
