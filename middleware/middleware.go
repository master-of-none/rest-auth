package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/master-of-none/rest-auth/utils"
)

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
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Could not validate the token",
			})
			return
		} else {
			if username, exists := claims["username"].(string); exists {
				ctx.Set("username", username)
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "Username does not exist",
				})
				return
			}
		}

		// ctx.Set("username", claims["username"].(string))

		ctx.Next()
	}
}
