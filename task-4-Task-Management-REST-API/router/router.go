package router

import (
	"github.com/gin-gonic/gin"
	"github.com/philipos/api/controllers"
)

func TaskRouters (router *gin.Engine) {
	router.GET("/books", controllers.GetAllTasks)
	router.GET("/book/:id", controllers.GetTask)
	router.POST("/create", controllers.CreateTask)
	router.PUT("/update/:id", controllers.UpdateTask)
	router.DELETE("/delete/:id", controllers.DeleteTask)
}