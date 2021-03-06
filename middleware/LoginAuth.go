package middleware

import (
	"github.com/gin-gonic/gin"
	"project/utils"
)

func LoginAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		if token == "" {
			utils.Error(context, "未登录")
			context.Abort()
			return
		} else {
			err, data := utils.CheckJwt(token)
			if err {
				utils.Error(context, data)
				context.Abort()
				return
			} else {
				//数据 进入协程上下文
				context.Set("user", data)
			}
		}
		context.Next()
	}
}
