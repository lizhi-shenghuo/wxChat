package controller

import (
	"fmt"
	"getaway/dao/redis"
	"getaway/service"
	mater "github.com/silenceper/wechat/v2/officialaccount/material"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"net/http"
	//"github.com/silenceper/wechat/v2/officialaccount/message"
)

func ServeWechat(rw http.ResponseWriter, req *http.Request) {
	// 传入request和responseWriter
	// 微信起一个路由即可 controller层根据事件触发service对应消息返回
	// 传入request和responseWriter
	server := service.OfficialAccount.GetServer(req, rw)

	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//TODO
		//回复消息：演示回复用户发送的消息
		//msg.MsgType
		fmt.Printf("msg----->:%#v\n", msg)
		switch msg.Event {
		case message.EventSubscribe:
			con, err := redis.RedisClient.Get("con").Result()
			if err != nil {
				fmt.Printf("redis get con failed, error(%v)\n", err)
			}

			text := message.NewText(con)
			fmt.Println(text)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		}

		switch msg.Content {
		case "明白":
			m := service.OfficialAccount.GetMaterial()
			mediaID, _, err := m.AddMaterial(mater.MediaTypeImage, "./images/105.png")
			if err != nil {
				fmt.Printf("m.AddMaterial error , err=%v\n", err)
			}
			image := message.NewImage(mediaID)
			return &message.Reply{MsgType: message.MsgTypeImage, MsgData: image}
		default:
			text := message.NewText("机器人不知道你在说啥呢")
			fmt.Println(text)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		}
	})

	//处理消息接收以及回复
	server.Serve()

	//发送回复的消息
	server.Send()
}
