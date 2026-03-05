package main

import (
	"github.com/philipos/api/delivery/controllers"
	"github.com/philipos/api/delivery/router"
	"github.com/philipos/api/infrastructure"
	"github.com/philipos/api/repository"
	"github.com/philipos/api/usecase"
)

func main() {
	db := infrastructure.ConnectDB()
	taskRepo := repository.NewTaskRepository(db, "tasks")
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskController := controllers.NewTaskController(taskUsecase)

	userRepo := repository.NewUserRepository(db, "users")
	userUsecase := usecase.NewUserUsecase(userRepo)
	userController := controllers.NewUserController(userUsecase)

	r := router.TaskRouters(taskController, userController)
	r.Run("localhost:3000")
}
