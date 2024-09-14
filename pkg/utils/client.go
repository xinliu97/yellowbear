package utils

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	Client *mongo.Client
}

func NewMongoDBClient(uri string) (*MongoDBClient, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &MongoDBClient{Client: client}, nil
}

func (mc *MongoDBClient) GetCollection(dbName string, collName string) *mongo.Collection {
	coll := mc.Client.Database(dbName).Collection(collName)
	return coll
}

func (mc *MongoDBClient) Disconnect() error {
	if err := mc.Client.Disconnect(context.TODO()); err != nil {
		fmt.Println("[Disconnect] Failed to disconnect to mongodb client.", err)
		return err
	}
	return nil
}
