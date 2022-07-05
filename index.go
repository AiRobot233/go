package main

import (
	"github.com/gin-gonic/gin"
	"project/controller"
	"project/middleware"
	"project/validate"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/someGet", func(context *gin.Context) {
			bol := validate.LoginValidate(context)
			if bol {
				controller.Login(context)
			}
		})
		v1.POST("/redis", func(context *gin.Context) {
			controller.Redis()
		})
		v1.GET("/jwt", func(context *gin.Context) {
			controller.Jwt(context)
		})
		//新增数据
		v1.POST("/save", func(context *gin.Context) {
			controller.SqlSave(context)
		})
		//删除数据
		v1.DELETE("/del", func(context *gin.Context) {
			controller.SqlDel()
		})
	}

	v2 := router.Group("/v2")
	{
		v2.Use(middleware.LoginAuth()).POST("/check_jwt", func(context *gin.Context) {
			controller.A(context)
		})
	}
	_ = router.Run(":8080")
}
