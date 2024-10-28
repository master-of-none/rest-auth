package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/master-of-none/rest-auth/controller"
	"github.com/master-of-none/rest-auth/middleware"
)

func AdminRoutes(r *gin.Engine) {
	adminRoute := r.Group("/admin")
	adminRoute.Use(middleware.AuthMiddleWare(), middleware.AdminMiddleware())
	{
		adminRoute.DELETE("/delete/:username", controller.DeleteUser)
	}

}
