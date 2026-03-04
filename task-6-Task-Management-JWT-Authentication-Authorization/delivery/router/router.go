package router

import (
	"github.com/gin-gonic/gin"
	"github.com/philipos/api/delivery/controllers"
)

func TaskRouters(controller *controllers.TaskController) *gin.Engine {
	// router.POST("/register", controllers.Register)
	// router.POST("/login", controllers.Login)
	router := gin.Default()
	authRouter := router.Group("/tasks")
	// authRouter.Use(middleware.AuthMiddleware())

	{
		authRouter.GET("", controller.FetchAll)
		authRouter.GET("/:id", controller.GetByID)
		authRouter.POST("", controller.Create)
		authRouter.PUT("/:id", controller.Update)
		// authRouter.DELETE("/:id", middleware.RoleMiddleware("admin"), controller.Delete)
		authRouter.DELETE("/:id", controller.Delete)

	}

	return router
}
