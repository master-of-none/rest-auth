package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
	ctx.SetCookie("Authorization", "", -1, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logout Successful",
	})
}
