package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"project/config"
	"project/model"
	"project/utils"
	"time"
)

var where [][]string
var m = model.SftLawMainBody{}

func Login(context *gin.Context) {
	password := context.PostForm("password")
	username := context.PostForm("username")
	where = [][]string{
		{"pid", "=", "1"},
	}
	data := m.First(where, "id,name")
	data["password"] = password
	data["username"] = username
	utils.Success(context, data)
}

func Redis() {
	rdb := config.BuildRedis()
	ctx := context.Background()
	sth, _ := time.ParseDuration("1h")
	res := rdb.Set(ctx, "key", "asdadada", sth)
	fmt.Println(res)
}
