package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}

	// Optional: Ping to test connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB Ping Error:", err)
	}

	MongoDB = client
	log.Println("Connected to MongoDB!")
}

func GetCollection(collectionName string) *mongo.Collection {
	return MongoDB.Database("toDodb").Collection(collectionName)
}
