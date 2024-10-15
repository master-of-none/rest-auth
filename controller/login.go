package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/master-of-none/rest-auth/databases"
	"github.com/master-of-none/rest-auth/models"
	"github.com/master-of-none/rest-auth/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func LoginCheck(ctx *gin.Context) {
	var loginRequest models.User
	var user models.User

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}
	var MongoClient *mongo.Client = databases.ConnectDB(ctx)
	collection := MongoClient.Database("users").Collection("user_info")
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//* Have to do Seperate Username and password check because in DB password is hashed.
	//* So first retreiving the username and checking, then retrieving the password and checking

	var filter bson.M
	if loginRequest.Username != "" {
		filter = bson.M{"username": loginRequest.Username}
	} else if loginRequest.Email != "" {
		filter = bson.M{"email": loginRequest.Email}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Username or Email must be provided",
		})
		return
	}
	//? For username check
	err := collection.FindOne(ctxMongo, filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Invalid Username or Email",
				"details": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error retreiving user data",
			})
		}
		return
	}

	//? This is for the password check
	passVal := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !passVal {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid password",
		})
		return
	}
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to generate the token",
			"details": err.Error(),
		})
		return
	}
	refreshToken, err := utils.GenerateRefreshToken(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to generate Refresh Token",
			"details": err.Error(),
		})
		return
	}

	//* Store Tokens in Cookie
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", token, 3600*2, "", "", false, true) // 2 hours

	//? Store Refresh Token
	ctx.SetCookie("RefreshToken", refreshToken, 3600*24*7, "", "", false, true) // 7 days
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Login Successful",
		"token":        token,
		"refreshToken": refreshToken,
	})
	//! Link: https://chatgpt.com/share/670c5b50-b1f0-8009-a430-ee84a5fc0698
}

func RefreshToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("RefreshToken")

	if err != nil || refreshToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "No Refresh Token provided",
			"details": err.Error(),
		})
		return
	}

	// Validate Refresh Token
	token, err := utils.ValidateToken(refreshToken)
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Invalid or expired Refresh Token",
			"details": err.Error(),
		})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Could not validate the token",
		})
		return
	}
	username := claims["username"].(string)

	newToken, err := utils.GenerateNewAccessToken(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Cannot generate new token",
			"details": err.Error(),
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", newToken, 3600*2, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"username":       username,
		"message":        "New Access token generated successfully",
		"newAccessToken": newToken,
	})
}
