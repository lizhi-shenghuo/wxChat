package controller

import (
	"fmt"
	mater "github.com/silenceper/wechat/v2/officialaccount/material"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"net/http"
)

// 回复和用户一样的话
func SendImgMsg(rw http.ResponseWriter, req *http.Request) {

	// 传入request和responseWriter
	server := officialAccount.GetServer(req, rw)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//TODO
		//回复图片：演示回复用户发送的消息
		println(msg.Content)
		m := officialAccount.GetMaterial()
		mediaID, url, err := m.AddMaterial(mater.MediaTypeImage, "./images/105.png")
		if err != nil {
			panic(err)
		}
		println("======", mediaID)
		fmt.Println(url)
		image := message.NewImage(mediaID)

		return &message.Reply{MsgType: message.MsgTypeImage, MsgData: image}
	})

	//处理消息接收以及回复
	server.Serve()

	//发送回复的消息
	server.Send()

}
