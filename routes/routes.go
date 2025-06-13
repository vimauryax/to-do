package routes

import (
	"github.com/Samratakgec/to-do-go-api/controller"
	"github.com/Samratakgec/to-do-go-api/middleware"
	"github.com/gin-gonic/gin"
)

func TaskRoutes (router *gin.Engine) {
	task := router.Group("/task")
	task.Use(middleware.JWTAuthMiddleware())
	{
		task.POST("/new",controller.CreateTask)
		task.GET("/get",controller.GetTaskById)
		task.PATCH("/update",controller.UpdateTask)
		task.DELETE("/delete",controller.DeleteTaskByID)
	}
}