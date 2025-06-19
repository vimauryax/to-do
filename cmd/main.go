package main

import (
	"github.com/Samratakgec/to-do-go-api/config"
	"github.com/Samratakgec/to-do-go-api/controller"
	cronjobs "github.com/Samratakgec/to-do-go-api/cron-jobs"
	"github.com/Samratakgec/to-do-go-api/routes"
	"github.com/Samratakgec/to-do-go-api/services"
	"github.com/gin-gonic/gin"
)

func main() {

	// connect with db
	config.ConnectDB()

	config.InitializeRedis()

	// initialize provider
	var taskService services.TaskService
	controller.InitSetterTaskService(taskService)

	// creating router
	r := gin.Default()

	// register routes
	routes.TaskRoutes(r)

	// call cron fn
	cronjobs.UseCron()

	// Start server
	r.Run(":8080")

}
