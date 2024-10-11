package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	// * Running on Port 8080
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the Environment File")
	}
	port := os.Getenv("PORT")
	r.Run(port)
}
