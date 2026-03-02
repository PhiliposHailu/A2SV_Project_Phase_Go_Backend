package router

import (
	"github.com/gin-gonic/gin"
	"github.com/philipos/api/controllers"
)

func TaskRouters (router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/tasks/:id", controllers.GetTask)
	router.POST("/tasks", controllers.CreateTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
}