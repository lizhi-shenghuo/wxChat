package main

//https://dashboard.ngrok.com/get-started/setup
//https://silenceper.com/wechat/officialaccount
//https://github.com/gowechat/example

import (
	"fmt"
	"getaway/config"
	"getaway/controller"
	"getaway/dao/mysql"
	"getaway/dao/redis"
	"getaway/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	err := redis.RsInit()
	if err != nil {
		fmt.Printf("connect redis error , err=%v\n", err)
		return
	}
	// 实例化wx对象
	gCfg := config.GetConfig()
	service.InitWechat(gCfg)
	err = mysql.Init(gCfg)
	if err != nil {
		log.Printf("connect redis error , err=%v\n", err)
		return
	}
	go func() {
		r := gin.Default()
		r.PUT("api/v1/UpdateSubScribeMsg/:data", controller.UpdatesubScribeMsg)
		r.PUT("api/v1/UpdateGroupQrCode/:data", controller.UpdateGroupQrCode)
		r.PUT("api/v1/UpdateCustomerMsg/:data", controller.UpdateCustomerMsg)
		r.PUT("api/v1/UpdateCustomerImgMsg/:data", controller.UpdateCustomerImgMsg)
		r.PUT("api/v1/UpdateSendTimer/:data", controller.UpdateSendTimer)

		r.Run(":6666")
	}()

	http.HandleFunc("/", controller.ServeWechat)
	fmt.Println("wechat server listener at", ":8090")
	err = http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Printf("start server error , err=%v\n", err)
		return
	}
}
