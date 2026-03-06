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
	passwordService := infrastructure.NewPasswordService()
	jwtService:= infrastructure.NewJWTService("mySecretePassword")

	userUsecase := usecase.NewUserUsecase(userRepo, passwordService, jwtService)

	userController := controllers.NewUserController(userUsecase)

	r := router.TaskRouters(taskController, userController, jwtService)
	r.Run("localhost:3000")
}
