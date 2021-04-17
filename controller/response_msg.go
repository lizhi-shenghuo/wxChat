package controller

import (
	"fmt"
	"net/http"

	"github.com/silenceper/wechat/v2/officialaccount/message"
)

func ServeWechat(rw http.ResponseWriter, req *http.Request) {

	// 传入request和responseWriter
	server := officialAccount.GetServer(req, rw)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//TODO
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		fmt.Println(text)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	server.Serve()

	//发送回复的消息
	server.Send()

}
