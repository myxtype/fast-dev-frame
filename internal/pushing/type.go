package pushing

import (
	"encoding/json"
)

type Request struct {
	Token  string   `json:"token"`  // 用户Token
	Topics []string `json:"topics"` // 订阅主题列表
}

type TopicResponse struct {
	Topic string      `json:"topic"` // 主题
	Data  interface{} `json:"data"`  // 数据
}

func (m TopicResponse) Bytes() []byte {
	b, _ := json.Marshal(m)
	return b
}
