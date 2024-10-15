package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/master-of-none/rest-auth/utils"
)

//! TODO - Implement MiddleWare

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//* Get the JWT Token
		tokenString, err := ctx.Cookie("Authorization")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired Auth Token",
			})
			ctx.Abort()
			return
		}
		token, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			fmt.Println(err.Error())
			refreshTokenString, refreshErr := ctx.Cookie("RefreshToken")
			if refreshErr != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error":   "Invalid or expired Refresh token",
					"details": err.Error(),
				})
				ctx.Abort()
				return
			}

			refreshTokenString, err := ctx.Cookie("RefreshToken")
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": "Invalid Token format",
				})
				ctx.Abort()
				return
			}
			refreshToken, err := utils.ValidateToken(refreshTokenString)
			if err != nil || !refreshToken.Valid {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error":   "Invalid or expired Refresh token",
					"details": err.Error(),
				})
				ctx.Abort()
				return
			}

			newToken, newTokenErr := utils.GenerateNewAccessToken(refreshToken)
			if newTokenErr != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Failed to generate new access Token",
					"message": newTokenErr.Error(),
				})
				ctx.Abort()
				return
			}
			ctx.SetCookie("Authorization", newToken, 3600, "", "", false, true)
			fmt.Println("New Access Token set since old expired")
		}
		ctx.Next()
	}
}
