package main

import (
	"github.com/gin-gonic/gin"
	"project/controller"
	"project/middleware"
	"project/validate"
)

func main() {
	router := gin.Default()
	router.POST("/someGet", func(context *gin.Context) {
		bol := validate.LoginValidate(context)
		if bol {
			controller.Login(context)
		}
	})
	router.POST("/redis", func(context *gin.Context) {
		controller.Redis()
	})
	router.GET("/jwt", func(context *gin.Context) {
		controller.Jwt(context)
	})
	router.Use(middleware.LoginAuth()).POST("/check_jwt", func(context *gin.Context) {
		controller.A(context)
	})
	_ = router.Run(":8080")
}
