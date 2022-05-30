package middleware

import (
	"github.com/gin-gonic/gin"
)

func LoginAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}
