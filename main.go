package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/master-of-none/rest-auth/config"
	"github.com/master-of-none/rest-auth/routes"
)

func main() {
	//* Load ENV file
	config.LoadEnv()
	r := gin.Default()

	//! 404 Error Handler
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "404 not found",
		})
	})
	routes.RegisterRoutes(r)

	//! TODO disconnect DB Code
	port := os.Getenv("PORT")
	r.Run(port)
}
