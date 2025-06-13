package models

//import "time"

type User struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Task struct {
	Task_id        string    `json:"task_id" bson:"task_id"`
	Title          string    `json:"title" binding:"required" bson:"title"`
	Desc           string    `json:"desc" bson:"desc"`
	TimestampBegin string `json:"timestampBegin" binding:"required" bson:"TimestampBegin"`
	TimestampEnd   string `json:"timestampEnd" binding:"required" bson:"timestampEnd"`
	Status         string    `json:"status" binding:"required,oneof=pending done outdated" bson:"status"`
}


// validator one-of
// validator pkg
