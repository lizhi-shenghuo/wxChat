package main

//https://dashboard.ngrok.com/get-started/setup
//https://silenceper.com/wechat/officialaccount
//https://github.com/gowechat/example

import (
	"fmt"
	"getaway/config"
	"getaway/controller"
	"getaway/dao/redis"
	"net/http"
)

func main() {
	err := redis.RsInit()
	if err != nil {
		fmt.Printf("connect redis error , err=%v\n", err)
	}
	// 实例化wx对象
	gCfg := config.GetConfig()
	controller.InitWechat(gCfg)

	// 路由太少 没有抽离
	//http.HandleFunc("/", controller.SendRepeatMsg)
	http.HandleFunc("/", controller.SendImgMsg)

	fmt.Println("wechat server listener at", ":8082")
	err = http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Printf("start server error , err=%v\n", err)
	}
}
