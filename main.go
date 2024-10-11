package main

import (
	"github.com/gin-gonic/gin"
	"github.com/master-of-none/rest-auth/routes"
)

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run()
}
