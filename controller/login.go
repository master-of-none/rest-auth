package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/master-of-none/rest-auth/databases"
	"github.com/master-of-none/rest-auth/models"
	"github.com/master-of-none/rest-auth/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func LoginCheck(ctx *gin.Context) {
	//! REDOING LOGIN wiht JWT - extra features mentioned in the end
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

	//? For username check
	err := collection.FindOne(ctxMongo, bson.M{"username": loginRequest.Username}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Invalid Username",
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

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login Successful",
		"token":   token,
	})

	//! TODO Middleware
	//! TODO Store token in DB
	//! TODO Refresh token
	//! Link: https://chatgpt.com/share/670c5b50-b1f0-8009-a430-ee84a5fc0698
}
