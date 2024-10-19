package databases

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/master-of-none/rest-auth/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetNextSequence(sequenceName string, ctx *gin.Context) (int, error) {
	var MongoClient *mongo.Client = ConnectDB(ctx)
	collection := MongoClient.Database("users").Collection("counters")
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": sequenceName}
	update := bson.M{"$inc": bson.M{"sequence_value": 1}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var result models.PostResult
	err := collection.FindOneAndUpdate(ctxMongo, filter, update, opts).Decode(&result)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error in updating the sequence value",
			"details": err.Error(),
		})
		return 0, err
	}
	return result.SequenceValue, nil
}
