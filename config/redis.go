package config

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Samratakgec/to-do-go-api/models"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
)

var RedisClient *redis.Client

func InitializeRedis() error {
	fmt.Println("start")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if err := syncMongoToRedis(); err != nil {
		fmt.Println("redis not activated")
		return err
	}
	return nil
}

func syncMongoToRedis() error {
	collection := GetCollection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	// Fetch all documents from MongoDB
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Mongo find error:", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var Task models.Task
		if err := cursor.Decode(&Task); err != nil {
			fmt.Println("Error decoding task:", err)
			continue
		}
		// Convert Task struct to JSON
		TaskJson, err := json.Marshal(Task)
		if err != nil {
			fmt.Println("Error marshaling task to JSON:", err)
			continue
		}
		err4 := RedisClient.Set(ctx, Task.Task_id, TaskJson, 5*time.Minute).Err()
		if err4 != nil {
			fmt.Println("not inserted in redis")
			return err4
		}
	}
	return nil

}
