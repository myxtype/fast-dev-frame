package i18n

// 全局i18n
var def *I18n

// 创建全局默认的i18n
func init() {
	def = NewI18n(map[string]map[string]string{
		"zh_CN": {
			"请求参数错误": "请求参数错误",
		},
		"zh_TW": {
			"请求参数错误": "請求參數錯誤",
		},
		"en_US": {
			"请求参数错误": "Request parameter error",
		},
		"ja_JP": {
			"请求参数错误": "要求パラメータエラー",
		},
		"ko_KR": {
			"请求参数错误": "요청 매개 변수 오류",
		},
	})
}

// 通过全局的i18n获取翻译
func Get(lang, key string) string {
	return def.Get(lang, key)
}
