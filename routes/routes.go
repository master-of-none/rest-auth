package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/master-of-none/rest-auth/controller"
	"github.com/master-of-none/rest-auth/databases"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello Rest World",
		})
	})

	r.POST("/login", controller.LoginCheck)
	//! TODO
	r.GET("/dbcheck", func(ctx *gin.Context) {
		client := databases.ConnectDB(ctx)

		if client == nil {
			//? Already an error is being sent
			return
		}

		ctx.JSON(200, gin.H{
			"message": "Database has been connected Successfully",
		})
	})
	// r.POST()
}
