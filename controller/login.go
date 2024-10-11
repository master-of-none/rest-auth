package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/master-of-none/rest-auth/models"
)

func LoginCheck(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON provided",
		})
	}

	//? Temporaroly check login without implementing DB
	if user.Username == "admin" && user.Password == "password" {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Login Successful",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Response",
		})
	}
}
