package services

import (
	"context"
	"errors"
	"time"

	"github.com/Samratakgec/to-do-go-api/config"
	"github.com/Samratakgec/to-do-go-api/helpers"
	"github.com/Samratakgec/to-do-go-api/models"
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

	_, err2 := collection.InsertOne(ctx, task)

	return err2
}

func GetTaskById(task_id string) (*models.Task, error) {
	collection := config.GetCollection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

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

	_, err := collection.DeleteOne(ctx, bson.M{"task_id": task_id})
	return err
}
