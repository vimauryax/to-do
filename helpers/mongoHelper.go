package helpers

import (
	"context"
	//"fmt"
	"strconv"
	"time"

	"github.com/Samratakgec/to-do-go-api/config"
	"go.mongodb.org/mongo-driver/bson"
)

func InitializeDb() (string, error) {
	return "sucessfully initialized", nil
}

func GetNextTaskId() (string, error) {

	collection := config.GetCollection("counter")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	filter := bson.M{"name": "task_id"}
	fetched_data := collection.FindOne(ctx, filter)
	//fmt.Println("uyghgugg ",fetched_data)

	var result struct {
		Seq int
	}
	err1 := fetched_data.Decode(&result)

	if err1 != nil {
		return "", err1
	}

	seq := result.Seq
	//fmt.Println("tfyftygy ",seq)
	seq++
	//fmt.Println(seq)
	update := bson.M{
		"$set": bson.M{
			"seq": seq,
		},
	}
	_, err2 := collection.UpdateOne(ctx, filter, update)

	return strconv.Itoa(seq), err2
}
