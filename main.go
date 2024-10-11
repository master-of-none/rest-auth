package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/master-of-none/rest-auth/routes"
)

func main() {
	r := gin.Default()

	//! 404 Error Handler
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "404 not found",
		})
	})
	routes.RegisterRoutes(r)

	r.Run()
}
