package databases

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoClient *mongo.Client

func ConnectDB(c *gin.Context) *mongo.Client {
	uri := os.Getenv("MONGO_URI")

	opts := options.Client().ApplyURI(uri)

	// Create a context with a timeout (optional but recommended)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connecting client and connecting to the server
	client, err := mongo.Connect(opts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to connect to the Database",
			"details": err.Error(),
		})
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			fmt.Printf("Error disconnecting from database: %v\n", err)
		}
	}()
	var result bson.M
	if err := client.Database("users").RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to ping the database",
			"details": err.Error(),
		})
	}
	MongoClient = client
	return MongoClient
}
