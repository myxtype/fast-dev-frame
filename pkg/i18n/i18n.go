package i18n

// 简单的i18n
type I18n struct {
	// lang:string => key:string => text:string
	data map[string]map[string]string
}

func NewI18n(data map[string]map[string]string) *I18n {
	return &I18n{data: data}
}

// 设置某一个语言
func (i *I18n) Set(lang, key, text string) {
	i.data[lang][key] = text
}

// 通过语言标识和Key来获取对应翻译
// 如果不存在直接原样返回
func (i *I18n) Get(lang, key string) string {
	if l, ok := i.data[lang]; ok {
		if v, ok := l[key]; ok {
			return v
		}
	}
	return key
}
