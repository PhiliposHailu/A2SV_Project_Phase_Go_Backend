package router

import (
	"github.com/gin-gonic/gin"
	"github.com/philipos/api/delivery/controllers"
	"github.com/philipos/api/middleware"
	"github.com/philipos/api/domain"

)

func TaskRouters(taskController *controllers.TaskController, userController *controllers.UserController, jwts domain.JWTService) *gin.Engine {
	router := gin.Default()

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	authRouter := router.Group("/tasks")
	authRouter.Use(middleware.AuthMiddleware(jwts))

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
