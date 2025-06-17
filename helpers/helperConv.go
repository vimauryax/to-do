package helpers

import (
	//"fmt"
	"strconv"
	"time"

	"github.com/Samratakgec/to-do-go-api/models"
)

func ConvertPayloadToTask(taskPayload models.TaskPayload) models.Task {
	loc, _ := time.LoadLocation("Asia/Kolkata")

	tBegin := taskPayload.TimestampBegin.In(loc).Unix()
	tEnd := taskPayload.TimestampEnd.In(loc).Unix()

	return models.Task{
		Title:          taskPayload.Title,
		Desc:           taskPayload.Desc,
		TimestampBegin: strconv.FormatInt(tBegin, 10),
		TimestampEnd:   strconv.FormatInt(tEnd, 10),
		Status:         taskPayload.Status,
	}
}
