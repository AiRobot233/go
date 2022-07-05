package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project/model"
	"project/utils"
)

func Jwt(context *gin.Context) {
	var data map[string]interface{}     //定义map
	data = make(map[string]interface{}) //初始化map
	data["uid"] = 1
	data["user_name"] = "hhh"
	token := utils.GetJwt(data, 21)
	utils.Success(context, token)
}

func A(context *gin.Context) {
	data, _ := context.Get("user")
	utils.Success(context, data)
}

func SqlSave(context *gin.Context) {
	var m = model.SftLawMainBody{}
	m.Name = "aaaaaa"
	m.Type = "机关单位"
	m.LawCategoryCode = "01"
	fmt.Println(m)
	res := m.Save(m)
	utils.Success(context, res)
}

func SqlDel() {
	var m = model.SftLawCategory{}
	m.Code = "01"
	res := m.Delete(m)
	fmt.Println(res)
}
