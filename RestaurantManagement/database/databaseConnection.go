package database

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectionDatabase() *mongo.Client {
	MongoDb_Url := "mongodb://localhost:27017"
	fmt.Println("MongoDB URI",MongoDb_Url)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb_Url))

	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second )

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	fmt.Println("Connected to MongoDB successfully")


	return client
}

var Client *mongo.Client = ConnectionDatabase()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("restaurant").Collection(collectionName)

	if collection == nil {
		log.Fatal("Error getting collection:", collectionName)
	}

	fmt.Println("Collection", collectionName, "opened successfully")
	return collection
}