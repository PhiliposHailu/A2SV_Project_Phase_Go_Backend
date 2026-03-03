package router

import (
	"github.com/gin-gonic/gin"
	"github.com/philipos/api/delivery/controllers"
	"github.com/philipos/api/middleware"
)

func TaskRouters(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	authRouter := router.Group("/tasks")
	authRouter.Use(middleware.AuthMiddleware())

	{
		authRouter.GET("", controllers.GetAllTasks)
		authRouter.GET("/:id", controllers.GetTask)
		authRouter.POST("", controllers.CreateTask)
		authRouter.PUT("/:id", controllers.UpdateTask)
		authRouter.DELETE("/:id", middleware.RoleMiddleware("admin"), controllers.DeleteTask)
	}
}
