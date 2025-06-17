package cronjobs

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/Samratakgec/to-do-go-api/config"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func checkForOutdated() {
	var Chan = make(chan string, 100)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	fmt.Println("is it running??")
	collection := config.GetCollection("task")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"status": "pending"})

	if err != nil && err != mongo.ErrNoDocuments {
		fmt.Println("error in finding documents")
	}
	defer cursor.Close(ctx)

	// looping in cursor
	for cursor.Next(ctx) {

		var task bson.M
		if err := cursor.Decode(&task); err != nil {
			fmt.Println("error decoding task:", err)
			continue
		}
		wg.Add(1)

		// Launch each task check/update in its own goroutine
		go func(task bson.M, Chan chan<- string, wg *sync.WaitGroup) {
			defer wg.Done()
			TimestampEnd, ok := task["timestampEnd"].(string)
			if !ok {
				fmt.Println("timestampEnd not found or not a string")
				return
			}

			t_end, err := strconv.ParseInt(TimestampEnd, 10, 64)
			if err != nil {
				fmt.Println("error parsing timestampEnd:", err)
				return
			}

			t_now := time.Now().Unix()
			if t_end < t_now {
				id, ok := task["_id"].(primitive.ObjectID)
				if !ok {
					fmt.Println("invalid ID format")
					return
				}

				update := bson.M{"$set": bson.M{"status": "outdated"}}
				_, err := collection.UpdateByID(ctx, id, update)
				if err != nil {
					fmt.Println("failed to update task:", err)
				} else {
					fmt.Println("marked as outdated:", id.Hex())
				}
			}
			mutex.Lock()
			fmt.Println("Sending message to channel")
			Chan <- "a go-routine done!!"
			mutex.Unlock()

		}(task, Chan, &wg)
	}

	wg.Wait()

	fmt.Println("channel data : ", Chan, " ", len(Chan))
	for msg := range Chan {
		fmt.Println("yes")
		fmt.Println("exact channel data : ", msg)
	}

	close(Chan)

}

func UseCron() {
	fmt.Println("welcome to cron-job pkg!!")
	c := cron.New()
	c.AddFunc("@every 10s", checkForOutdated)
	c.Start()
}
