package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewEmptyFilter()  {
	
}

func MongoUpdate(coll *mongo.Collection, filter bson.D, content bson.D) (*mongo.UpdateResult, error) {
	update := bson.D{{Key: "$set", Value: content}}
	updateResult, err := coll.UpdateOne(context.TODO(), filter, update)
	return updateResult, err
}