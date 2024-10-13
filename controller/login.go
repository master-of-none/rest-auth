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
	//! REDOING LOGIN wiht JWT
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

	err := collection.FindOne(ctxMongo, bson.M{"username": loginRequest.Username, "password": loginRequest.Password}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid username or password",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error retreiving user data",
			})
		}
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate the token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login Successful",
		"token":   token,
	})
}
