package i18n

// 全局i18n
var Def = NewI18n()

// 通过全局的i18n获取翻译
func Get(lang, key string, args ...Options) string {
	return Def.Get(lang, key, args...)
}

func SetValue(lang, key, value string) {
	Def.SetValue(lang, key, value)
}

func SetLangValues(lang string, values map[string]string) {
	Def.SetLangValues(lang, values)
}
