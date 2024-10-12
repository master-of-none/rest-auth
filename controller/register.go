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

func RegisterUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON Format",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot hash the password",
		})
		return
	}
	checkPassword := utils.CheckPasswordHash(user.Password, hashedPassword)
	if checkPassword != true {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot hash the password",
		})
		return
	}
	user.Password = hashedPassword

	var MongoClient *mongo.Client = databases.ConnectDB(ctx)
	collection := MongoClient.Database("users").Collection("user_info")
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//* Check User exists ✅
	var existingUser models.User
	err = collection.FindOne(ctxMongo, bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		//? User exists
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "Username already exists",
		})
		return
	}
	if err != mongo.ErrNoDocuments {
		// Error is another error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error in checking the existing user",
			"details": err.Error(),
		})
		return
	}
	_, err = collection.InsertOne(ctxMongo, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create user in the database",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User Created Successfully",
	})
}
