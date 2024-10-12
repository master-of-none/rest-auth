package databases

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoClient *mongo.Client

func ConnectDB(c *gin.Context) *mongo.Client {
	uri := os.Getenv("MONGO_URI")

	opts := options.Client().ApplyURI(uri)

	// Create a context with a timeout (optional but recommended)
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connecting client and connecting to the server
	client, err := mongo.Connect(opts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to connect to the Database",
			"details": err.Error(),
		})
	}
	MongoClient = client
	return MongoClient
}

//! TODO Disconnect
