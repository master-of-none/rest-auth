package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

//! TODO - Implement MiddleWare

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//* Get the JWT Token
		// authHeader := ctx.GetHeader("Authorization")

		// if authHeader == "" {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{
		// 		"error": "Authorization Header is missing",
		// 	})
		// 	ctx.Abort()
		// 	return
		// }

		// tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		// if tokenString == authHeader {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{
		// 		"error": "Invalid Token Format",
		// 	})
		// 	ctx.Abort()
		// 	return
		// }
		// fmt.Println(tokenString)
		tokenString, err := ctx.Cookie("Authorization")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Token format",
			})
			ctx.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Invalid or expired token",
				"details": err.Error(),
			})
			ctx.Abort()
			return
		}

		//! TODO: Validate Refresh Token
		//!
		ctx.Next()
	}
}
