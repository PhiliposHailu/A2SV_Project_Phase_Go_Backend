package router

import (
	"github.com/gin-gonic/gin"
	"github.com/philipos/api/delivery/controllers"
	"github.com/philipos/api/middleware"
)

func TaskRouters(taskController *controllers.TaskController, userController *controllers.UserController) *gin.Engine {
	router := gin.Default()

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	authRouter := router.Group("/tasks")
	authRouter.Use(middleware.AuthMiddleware())

	{
		authRouter.GET("", taskController.FetchAll)
		authRouter.GET("/:id", taskController.GetByID)
		authRouter.POST("", taskController.Create)
		authRouter.PUT("/:id", taskController.Update)
		authRouter.DELETE("/:id", middleware.RoleMiddleware("admin"), taskController.Delete)
		// authRouter.DELETE("/:id", taskController.Delete)

	}

	return router
}
