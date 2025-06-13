package main

import (
	"github.com/Samratakgec/to-do-go-api/config"
	cronjobs "github.com/Samratakgec/to-do-go-api/cron-jobs"
	"github.com/Samratakgec/to-do-go-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// connect with db
	config.ConnectDB()

	// creating router
	r := gin.Default()

	// register routes
	routes.TaskRoutes(r)

	// call cron fn
	cronjobs.UseCron()

	// Start server
	r.Run(":8080")


}
