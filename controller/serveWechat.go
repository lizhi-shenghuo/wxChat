package controller

import (
	"getaway/service"
	"net/http"
	//"github.com/silenceper/wechat/v2/officialaccount/message"
)

func ServeWechat(rw http.ResponseWriter, req *http.Request) {
	// 传入request和responseWriter
	// 微信起一个路由即可 controller层根据事件触发service对应消息返回

	service.SendImgMsg(rw, req)
}
