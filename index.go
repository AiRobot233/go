package main

import (
	"github.com/gin-gonic/gin"
	"project/controller"
	"project/middleware"
	"project/validate"
)

func main() {
	router := gin.Default()
	router.Use(middleware.LoginAuth()).POST("/someGet", func(context *gin.Context) {
		bol := validate.LoginValidate(context)
		if bol {
			controller.Login(context)
		}
	})
	_ = router.Run(":8080")
}
