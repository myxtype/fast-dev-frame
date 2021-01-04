package i18n

import "testing"

func TestNewI18n(t *testing.T) {
	i := NewI18n()
	i.SetValue("zh_CN", "title", "{{name}}，你好啊！")
	i.SetValue("zh_CN", "title2", "{{name}}，{{name}}，你好吗！！")

	t.Log(i.Get("zh_CN", "title", Options{"name": "阿狸"}))
	t.Log(i.Get("zh_CN", "title", Options{"name": 33442}))
	t.Log(i.Get("zh_CN", "title2", Options{"name": "阿狸"}))
	t.Log(i.Get("zh_CN", "title2"))
}
