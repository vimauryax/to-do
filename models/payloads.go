package models

import "time"

type TaskPayload struct {
	Title          string    `json:"title" binding:"required" bson:"title"`
	Desc           string    `json:"desc" bson:"desc"`
	TimestampBegin time.Time `json:"timestampBegin" binding:"required" bson:"TimestampBegin"`
	TimestampEnd   time.Time `json:"timestampEnd" binding:"required" bson:"timestampEnd"`
	Status         string    `json:"status" binding:"required,oneof=pending done outdated" bson:"status"`
}

type TaskUpdatePayload struct {
	Title        *string    `json:"title,omitempty"`
	Desc         *string    `json:"desc,omitempty"`
	Status       *string    `json:"status,omitempty"`
	TimestampEnd *time.Time `json:"timestampEnd,omitempty"`
}
