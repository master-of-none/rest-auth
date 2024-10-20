package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/master-of-none/rest-auth/controller"
	"github.com/master-of-none/rest-auth/databases"
	"github.com/master-of-none/rest-auth/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello Rest World",
		})
	})

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
	//? User Register ✅
	r.POST("/register", controller.RegisterUser)
	//* Middleware route
	protectedRoute := r.Group("/protected")
	protectedRoute.Use(middleware.AuthMiddleWare())
	{
		protectedRoute.GET("/dashboard", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message":  "Welcome to the protected Dashboard",
				"username": ctx.GetString("username"),
			})
		})
	}
	r.POST("/login", controller.LoginCheck)
	r.POST("/logout", controller.Logout)
	r.POST("/refreshToken", controller.RefreshToken)
	PostRoutes(r)
}
