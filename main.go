package main

import (
	"backend/controllers"
	"backend/db"
	"backend/domain/repository"
	"backend/domain/service"
	"backend/router"
	"backend/usecases"
)

func main() {
	database := db.ConnectDatabase()

	// repository
	userRepository := repository.NewMySQLUserRepository(database)
	organizationRepository := repository.NewMySQLOrganizationRepository(database)

	// service
	tokenService := service.NewTokenService()

	// usecase
	userUsecase := usecases.NewUserUsecase(userRepository, tokenService)
	organizationUsecase := usecases.NewOrganizationUsecase(organizationRepository)

	// controller
	userController := controllers.NewUserController(userUsecase)
	organizationController := controllers.NewOrganizationController(organizationUsecase)

	r := router.SetUpRouter()

	user := r.Group("/user")
	{
		user.POST("/signup", userController.SignUp)
		user.POST("/signin", userController.SignIn)
		user.GET("/verify", userController.VerifyToken)
	}

	organization := r.Group("/organization")
	{
		
	}

	r.Run(":8080")
}