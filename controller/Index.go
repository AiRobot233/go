package controller

import (
	"github.com/gin-gonic/gin"
	"project/utils"
)

func Jwt(context *gin.Context) {
	var data map[string]interface{}     //定义map
	data = make(map[string]interface{}) //初始化map
	data["uid"] = 1
	data["user_name"] = "hhh"
	token := utils.GetJwt(data, 21)
	utils.Success(context, "成功", token)
}

func A(context *gin.Context) {
	data, _ := context.Get("user")
	utils.Success(context, "", data)
}