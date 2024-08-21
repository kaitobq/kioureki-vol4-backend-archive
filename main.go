package main

import (
	"backend/controllers"
	"backend/db"
	"backend/domain/repository"
	"backend/router"
	"backend/usecases"
)

func main() {
	database := db.ConnectDatabase()

	// repository
	userRepository := repository.NewMySQLUserRepository(database)

	// usecase
	userUsecase := usecases.NewUserUsecase(userRepository)

	// controller
	userController := controllers.NewUserController(userUsecase)

	r := router.SetUpRouter()

	user := r.Group("/user")
	{
		user.POST("/signup", userController.SignUp)
		user.POST("/signin", userController.SignIn)
	}

	r.Run(":8080")
}