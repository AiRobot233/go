package controller

import (
	"github.com/gin-gonic/gin"
	"project/model"
	"project/utils"
)

var where [][]string
var m = model.SftLawMainBody{}

func Login(context *gin.Context) {
	where = [][]string{
		{"pid", "=", "1"},
	}
	data := m.First(where, "id,name")
	utils.Success(context, "", data)
}
