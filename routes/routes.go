package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/master-of-none/rest-auth/controller"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello Rest World",
		})
	})

	r.POST("/login", controller.LoginCheck)
}
