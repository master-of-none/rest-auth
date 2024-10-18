package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/master-of-none/rest-auth/controller"
	"github.com/master-of-none/rest-auth/middleware"
)

func PostRoutes(r *gin.Engine) {
	posts := r.Group("/posts")
	posts.Use(middleware.AuthMiddleWare())
	{
		posts.POST("/", controller.CreatePost)
		posts.GET("/", controller.GetPosts)
	}
}
