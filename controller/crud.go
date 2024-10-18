package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/master-of-none/rest-auth/databases"
	"github.com/master-of-none/rest-auth/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// ! Create Post
func CreatePost(ctx *gin.Context) {
	// ! TODO
}

// ! Get Posts
func GetPosts(ctx *gin.Context) {
	//! TODO
	var posts []models.Post
	var MongoClient *mongo.Client = databases.ConnectDB(ctx)

	collection := MongoClient.Database("users").Collection("posts")
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctxMongo, bson.D{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error in Retrieving the posts",
			"details": err.Error(),
		})
		return
	}
	defer cursor.Close(ctxMongo)

	for cursor.Next(ctxMongo) {
		var post models.Post
		if err := cursor.Decode(&post); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error in decoding the data",
				"details": err.Error(),
			})
			return
		}
		posts = append(posts, post)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

// ! Update posts
func UpdatePost(ctx *gin.Context) {
	//! TODO
}

// ! Delete Post
func DeletePost(ctx *gin.Context) {
	//! TODO
}
