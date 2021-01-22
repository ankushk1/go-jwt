package db

import (
	"context"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var collection *mongo.Collection

func GetDBCollection() (*mongo.Collection, error) {
	clientOptions :=options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	// fmt.Println("Connected to MongoDB!")
	collection := client.Database("GoLogin").Collection("users")
	return collection, nil
}