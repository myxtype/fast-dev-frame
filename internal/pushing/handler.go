package pushing

import (
	"frame/pkg/logger"
	"github.com/myxtype/go-webreal"
)

type Handler struct{}

func (b *Handler) OnConnect(client *webreal.Client) {
	logger.Sugar.Debugf("client(%v) connected", client.ID())

	// 自动订阅
	client.Subscribe("lucky_star")
}

func (b *Handler) OnClose(client *webreal.Client) {
	logger.Sugar.Debugf("client(%v) closed", client.ID())
}

func (b *Handler) OnMessage(client *webreal.Client, msg *webreal.Message) {
	switch msg.Type {
	case "subscribe":
		b.onSub(client, msg)
	case "unsubscribe":
		b.onUnsub(client, msg)
	}
}

// 订阅消息
func (b *Handler) onSub(client *webreal.Client, msg *webreal.Message) {
	var req Request
	if err := msg.ReadMessageData(&req); err != nil {
		return
	}

	for _, topic := range req.Topics {
		client.Subscribe(topic)
	}
}

// 取消订阅
func (b *Handler) onUnsub(client *webreal.Client, msg *webreal.Message) {
	var req Request
	if err := msg.ReadMessageData(&req); err != nil {
		return
	}

	for _, topic := range req.Topics {
		client.Unsubscribe(topic)
	}
}
