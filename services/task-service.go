package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Samratakgec/to-do-go-api/config"
	"github.com/Samratakgec/to-do-go-api/helpers"
	"github.com/Samratakgec/to-do-go-api/models"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
)

func CreateTask(task models.Task) error {
	collection := config.GetCollection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	task_id, err1 := helpers.GetNextTaskId()
	if err1 != nil {
		return err1
	}
	task.Task_id = task_id

	// Insert into MongoDB
	_, err2 := collection.InsertOne(ctx, task)
	if err2 != nil {
		return err2
	}

	// Marshal the task to JSON before caching in Redis
	taskJson, err3 := json.Marshal(task)
	if err3 != nil {
		fmt.Println("error in marshal")
		return err3
	}

	err4 := config.RedisClient.Set(ctx, task_id, taskJson, 5*time.Minute).Err()
	if err4 != nil {
		fmt.Println("not inserted in redis")
		return err4
	}
	fmt.Println("inserted in reddis")
	return nil
}

func GetTaskById(task_id string) (*models.Task, error) {
	collection := config.GetCollection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	data, err1 := config.RedisClient.Get(ctx, task_id).Result()
	if err1 == redis.Nil {
		fmt.Println("cache miss")
	} else if err1 != nil {
		fmt.Println("redis error:", err1)
	} else {
		fmt.Println("cache hit")
		var task models.Task
		err2 := json.Unmarshal([]byte(data), &task)
		if err2 != nil {
			fmt.Println("json unmarshal error:", err2)
		} else {
			return &task, nil
			// return or use task
		}

	}
	var task models.Task
	err := collection.FindOne(ctx, bson.M{"task_id": task_id}).Decode(&task)

	if err != nil {
		return nil, err
	}
	return &task, err
}

// func GetAllTask()  {
// collection := config.GetCollection("task")
// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

//		defer cancel()
//		collection.
//	}

func UpdateTaskByID(taskID string, update *models.TaskUpdatePayload) error {

	collection := config.GetCollection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err1 := config.RedisClient.Del(ctx, taskID).Err()
	if err1 != nil {
		fmt.Println("error in deleting redis")
		return err1
	}

	updateFields := bson.M{}

	if update.Title != nil {
		updateFields["title"] = *update.Title
	}
	if update.Desc != nil {
		updateFields["desc"] = *update.Desc
	}
	if update.Status != nil {
		updateFields["status"] = *update.Status
	}
	if update.TimestampEnd != nil {
		updateFields["timestampOut"] = *update.TimestampEnd
	}

	if len(updateFields) == 0 {
		return errors.New("no fields to update")
	}

	filter := bson.M{"task_id": taskID}
	updateDoc := bson.M{"$set": updateFields}

	_, err := collection.UpdateOne(ctx, filter, updateDoc)
	return err
}

func DeleteTaskByID(task_id string) error {
	collection := config.GetCollection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err1 := config.RedisClient.Del(ctx, task_id).Err()
	if err1 != nil {
		fmt.Println("error in deleting redis")
		return err1
	}

	_, err := collection.DeleteOne(ctx, bson.M{"task_id": task_id})
	return err
}
