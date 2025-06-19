package controller

import (
	// "time"

	// "github.com/Samratakgec/to-do-go-api/helpers"
	"fmt"
	//"strconv"
	"time"

	"github.com/Samratakgec/to-do-go-api/helpers"
	"github.com/Samratakgec/to-do-go-api/models"
	"github.com/Samratakgec/to-do-go-api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskServiceMongo models.TaskServiceMongo

func InitSetterTaskService(taskService services.TaskService) {
	taskServiceMongo = taskService
}

func CreateTask(c *gin.Context) { // 400, 500, 200
	var taskPayload models.TaskPayload
	var task models.Task

	if err := c.ShouldBindBodyWithJSON(&taskPayload); err != nil {
		c.JSON(400, gin.H{"error": "in-valid task payload"})
		return
	}
	loc, _ := time.LoadLocation("Asia/Kolkata")
	t_beg := taskPayload.TimestampBegin.In(loc).Unix()
	t_end := taskPayload.TimestampEnd.In(loc).Unix()
	t_now := time.Now().In(loc).Unix()
	fmt.Println("time beg : ", t_beg)
	fmt.Println("time end : ", t_end)
	fmt.Println("time now : ", t_now)

	fmt.Println(time.Now().In(loc).Format("2006-01-02 15:04"))
	if t_end < t_now {
		c.JSON(400, gin.H{"error": "invalid task payload: end time is before current time"})
		return
	}

	if t_beg > t_end {
		c.JSON(400, gin.H{"error": "invalid task payload: start time is after end time"})
		return
	}

	task = helpers.ConvertPayloadToTask(taskPayload)

	//err2 := services.CreateTask(task)
	err2 := taskServiceMongo.CreateTask(task)

	if err2 != nil {
		fmt.Println(err2)
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(200, gin.H{"success": "task added successfully"})
}

func GetTaskById(c *gin.Context) { // 400, 404, 500, 200
	task_id := c.Query("task_id")

	if task_id == "" {
		c.JSON(400, gin.H{"error": "no task_id provided in query params"})
		return
	}
	//var task models.Task
	task, err := taskServiceMongo.GetTaskById(task_id)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{"error": "no record found"})
			return
		} else {
			c.JSON(500, gin.H{"error": "internal server error"})
			return
		}
	}
	c.JSON(200, task)
}

func UpdateTask(c *gin.Context) { // 400, 404, 500, 200
	task_id := c.Query("task_id")

	var updatePayload models.TaskUpdatePayload
	if err := c.ShouldBindJSON(&updatePayload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("this should work", task_id)
	_, err1 := taskServiceMongo.GetTaskById(task_id)

	if err1 != nil {
		if err1 == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{"error": "no record found"})
			return
		} else {
			c.JSON(500, gin.H{"error": "internal server error"})
			return
		}
	}

	err := taskServiceMongo.UpdateTaskByID(task_id, &updatePayload)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "Task updated successfully"})
}

func DeleteTaskByID(c *gin.Context) {
	task_id := c.Query("task_id")

	if task_id == "" {
		c.JSON(400, gin.H{"error": "no task_id provided in query params"})
		return
	}

	_, err1 := taskServiceMongo.GetTaskById(task_id)
	if err1 != nil {
		if err1 == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{"error": "no record found"})
			return
		} else {
			c.JSON(500, gin.H{"error": "internal server error"})
			return
		}
	}

	err2 := taskServiceMongo.DeleteTaskByID(task_id)
	if err2 != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(200, gin.H{"success": "task deleted successfully"})
}
