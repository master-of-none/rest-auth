package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/master-of-none/rest-auth/databases"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func DeleteUser(ctx *gin.Context) {
	usernameDelete := ctx.Param("username")
	if usernameDelete == "admin" {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "cannot delete admin user",
		})
		return
	}

	var MongoClient *mongo.Client = databases.ConnectDB(ctx)
	collection := MongoClient.Database("users").Collection("user_info")
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctxMongo, bson.M{"username": usernameDelete})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Cannot delete the user",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
