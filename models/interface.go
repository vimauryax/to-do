package models

type TaskServiceMongo interface {
	CreateTask(task Task) error
	GetTaskById(task_id string) (*Task, error)
	UpdateTaskByID(taskID string, update *TaskUpdatePayload) error
	DeleteTaskByID(task_id string) error
}
