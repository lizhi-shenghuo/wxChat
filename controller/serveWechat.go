package controller

import (
	"fmt"
	"getaway/dao/redis"
	"getaway/service"
	mater "github.com/silenceper/wechat/v2/officialaccount/material"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"log"
	"net/http"
	"strconv"
	"time"
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
			subScribeMsg, err := redis.RedisClient.Get("subScribeMsg").Result()
			if err != nil {
				fmt.Printf("redis get subScribeMsg failed, error(%v)\n", err)
			}
			text := message.NewText(subScribeMsg)
			//fmt.Println(text)
			// 一分钟后发送文本
			sendTimer, err := redis.RedisClient.Get("sendTimer").Result()
			if err != nil {
				fmt.Printf("redis set sendTimer failed, error(%v)\n", err)
			}
			min, err := strconv.Atoi(sendTimer)
			if err != nil {
				log.Println("strconv.Atoi(sendTimer) error")
			}
			go time.AfterFunc(time.Duration(min)*time.Second, func() {
				mgr := message.NewMessageManager(service.OfficialAccount.GetContext())
				// 文本信息
				customerMsg, err := redis.RedisClient.Get("customerMsg").Result()
				if err != nil {
					fmt.Printf("redis get customerMsg failed, error(%v)\n", err)
				}
				textMsg := message.NewCustomerTextMessage(string(msg.FromUserName), customerMsg)
				err = mgr.Send(textMsg)
				if err != nil {
					log.Println("send text msg failed, err: ", err)
				}
				// 图片信息
				m := service.OfficialAccount.GetMaterial()
				customerImgMsg, err := redis.RedisClient.Get("customerImgMsg").Result()
				if err != nil {
					fmt.Printf("redis get customerImgMsg failed, error(%v)\n", err)
				}
				mediaID, _, err := m.AddMaterial(mater.MediaTypeImage, "./images/"+customerImgMsg)
				if err != nil {
					log.Println("add img msg failed, err: ", err)
				}
				imgMsg := message.NewCustomerImgMessage(string(msg.FromUserName), mediaID)
				err = mgr.Send(imgMsg)
				if err != nil {
					log.Println("send img msg failed, err: ", err)
				}
			})
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		}

		var image *message.Image
		switch msg.Content {
		case "rm":
			var text *message.Text
			_, err := redis.RedisClient.HDel("userImage", msg.GetOpenID()).Result()
			if err != nil {
				log.Printf("del %v redis cache failed, err: %v", msg.GetOpenID(), err)
				text = message.NewText("清除记录失败...")
			} else {
				text = message.NewText("清除记录成功...")
			}
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		case "明白":
			mediaIdCache, err := redis.RedisClient.HGet("userImage", msg.GetOpenID()).Result()
			m := service.OfficialAccount.GetMaterial()
			if err == redis.Nil {
				groupQrCode, err := redis.RedisClient.Get("groupQrCode").Result()
				if err != nil {
					fmt.Printf("redis get groupQrCode failed, error(%v)\n", err)
				}
				mediaID, _, err := m.AddMaterial(mater.MediaTypeImage, "./images/"+groupQrCode)
				if err != nil {
					fmt.Printf("m.AddMaterial error , err=%v\n", err)
				}
				image = message.NewImage(mediaID)
				// 添加缓存
				redis.RedisClient.HSet("userImage", msg.GetOpenID(), mediaID)
			} else if err != nil {
				log.Println("get redis hash cache failed, err: ", err)
			} else {
				image = message.NewImage(mediaIdCache)
			}

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
