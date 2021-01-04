package i18n

import (
	"github.com/spf13/cast"
	"strings"
)

// 简单的i18n
type I18n struct {
	data map[string]map[string]string
}

type Options map[string]interface{}

func NewI18n(data ...map[string]map[string]string) *I18n {
	if len(data) > 0 {
		return &I18n{data: data[0]}
	}
	return &I18n{
		data: map[string]map[string]string{},
	}
}

// 设置一个语言的KeyValue
func (i *I18n) SetValue(lang, key, value string) {
	if _, ok := i.data[lang]; !ok {
		i.data[lang] = map[string]string{}
	}
	i.data[lang][key] = value
}

// 设置整个语言的Value
func (i *I18n) SetLangValues(lang string, values map[string]string) {
	i.data[lang] = values
}

// 通过语言标识和Key来获取对应翻译
// 如果不存在直接原样返回key
func (i *I18n) Get(lang, key string, args ...Options) string {
	if values, ok := i.data[lang]; ok {
		if val, ok := values[key]; ok {
			if len(args) > 0 {
				return formatOptions(val, args[0])
			}
			return val
		}
	}
	return key
}

// 格式化参数
func formatOptions(val string, ops Options) string {
	for key, n := range ops {
		val = strings.Replace(val, "{{"+key+"}}", cast.ToString(n), -1)
	}
	return val
}
