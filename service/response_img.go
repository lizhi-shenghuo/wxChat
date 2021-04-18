package service

import (
	"fmt"
	mater "github.com/silenceper/wechat/v2/officialaccount/material"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/silenceper/wechat/v2/officialaccount/server"
)

// 回复和用户一样的话
func SendImgMsg(server *server.Server) *server.Server {

	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//TODO
		//回复图片：演示回复用户发送的消息
		println("========", msg.Event)
		m := OfficialAccount.GetMaterial()
		mediaID, url, err := m.AddMaterial(mater.MediaTypeImage, "./images/105.png")
		if err != nil {
			fmt.Printf("m.AddMaterial error , err=%v\n", err)
		}
		fmt.Println("=====", url)
		image := message.NewImage(mediaID)

		return &message.Reply{MsgType: message.MsgTypeImage, MsgData: image}
	})

	return server
}
