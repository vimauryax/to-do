package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

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

	DB = client
	log.Println("Connected to MongoDB!")
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database("toDodb").Collection(collectionName)
}
