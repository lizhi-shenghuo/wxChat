package controller

import (
	"getaway/dao/redis"
	"github.com/gin-gonic/gin"
	"log"
)

func UpdatesubScribeMsg(c *gin.Context) {
	data := c.Param("data")
	log.Println(data)
	result, err := redis.RedisClient.Set("subScribeMsg", data, -1).Result()
	if err != nil {
		log.Printf("subScribeMsg set error, err(%v)", err)
		ResponseErrorWithMsg(c, CodeServerBusy, "服务器错误")
	}
	ResponseSuccess(c, result)
}

//r.PUT("api/v1/UpdateGroupQrCode/:data", controller.UpdatesubScribeMsg)
//r.PUT("api/v1/UpdateCustomerMsg/:data", controller.UpdatesubScribeMsg)
//r.PUT("api/v1/UpdateCustomerImgMsg/:data", controller.UpdatesubScribeMsg)
func UpdateGroupQrCode(c *gin.Context) {
	data := c.Param("data")
	log.Println(data)
	result, err := redis.RedisClient.Set("groupQrCode", data, -1).Result()
	if err != nil {
		log.Printf("groupQrCode set error, err(%v)", err)
		ResponseErrorWithMsg(c, CodeServerBusy, "服务器错误")
	}
	ResponseSuccess(c, result)
}

func UpdateCustomerMsg(c *gin.Context) {
	data := c.Param("data")
	log.Println(data)
	result, err := redis.RedisClient.Set("customerMsg", data, -1).Result()
	if err != nil {
		log.Printf("customerMsg set error, err(%v)", err)
		ResponseErrorWithMsg(c, CodeServerBusy, "服务器错误")
	}
	ResponseSuccess(c, result)
}

func UpdateCustomerImgMsg(c *gin.Context) {
	data := c.Param("data")
	log.Println(data)
	result, err := redis.RedisClient.Set("customerImgMsg", data, -1).Result()
	if err != nil {
		log.Printf("customerImgMsg set error, err(%v)", err)
		ResponseErrorWithMsg(c, CodeServerBusy, "服务器错误")
	}
	ResponseSuccess(c, result)
}

func UpdateSendTimer(c *gin.Context) {
	data := c.Param("data")
	log.Println(data)
	result, err := redis.RedisClient.Set("sendTimer", data, -1).Result()
	if err != nil {
		log.Printf("sendTimer set error, err(%v)", err)
		ResponseErrorWithMsg(c, CodeServerBusy, "服务器错误")
	}
	ResponseSuccess(c, result)
}
