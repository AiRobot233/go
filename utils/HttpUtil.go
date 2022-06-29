package utils

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Success(context *gin.Context, msg string, data interface{}) {
	context.JSON(200, gin.H{
		"error":   0,
		"message": msg,
		"data":    data,
		"time":    time.Now().Unix(),
	})
}

func Error(context *gin.Context, msg interface{}) {
	context.JSON(400, gin.H{
		"error":   1,
		"message": msg,
		"data":    nil,
		"time":    time.Now().Unix(),
	})
}

//表单验证错误返回
func ValidateError(context *gin.Context, msg map[string]string) {
	context.JSON(412, gin.H{
		"error":   1,
		"message": msg,
		"data":    nil,
		"time":    time.Now().Unix(),
	})
}
