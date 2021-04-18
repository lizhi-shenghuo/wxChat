package service

import (
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/silenceper/wechat/v2/officialaccount/server"
)

// 回复和用户一样的话
func SendRepeatMsg(server *server.Server) *server.Server {

	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//TODO
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		fmt.Println(text)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})
	return server
}
