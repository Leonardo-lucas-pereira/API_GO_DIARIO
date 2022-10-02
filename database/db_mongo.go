package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func StartMongoDB() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://leonardo:senha123@localhost:27017"))
	if err != nil {
		panic(err)
	}

	collection = client.Database("diario").Collection("userPosts")
}

func GetDatabaseMongo() *mongo.Collection {
	return collection
}
