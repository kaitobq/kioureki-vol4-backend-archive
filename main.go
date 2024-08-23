package main

import (
	"backend/controllers"
	"backend/db"
	"backend/domain/repository"
	"backend/domain/service"
	"backend/middleware"
	"backend/router"
	"backend/usecases"
)

func main() {
	database := db.ConnectDatabase()

	// repository
	userRepository := repository.NewMySQLUserRepository(database)
	organizationRepository := repository.NewMySQLOrganizationRepository(database)
	userOrganizationInvitationRepository := repository.NewMySQLUserOrganizationInvitationRepository(database)
	userOrganizationMembershipRepository := repository.NewMySQLUserOrganizationMembershipRepository(database)

	// service
	tokenService := service.NewTokenService()

	// usecase
	userUsecase := usecases.NewUserUsecase(userRepository, tokenService, organizationRepository)
	organizationUsecase := usecases.NewOrganizationUsecase(organizationRepository, userOrganizationMembershipRepository, userRepository)
	userOrganizationInvitationUsecase := usecases.NewUserOrganizationInvitationUsecase(userOrganizationInvitationRepository, userOrganizationMembershipRepository, organizationRepository, userRepository)
	// userOrganizationMembershipUsecase := usecases.NewUserOrganizationMembershipUsecase(userOrganizationMembershipRepository)

	// controller
	userController := controllers.NewUserController(userUsecase)
	organizationController := controllers.NewOrganizationController(organizationUsecase)
	userOrganizationInvitationController := controllers.NewUserOrganizationInvitationController(userOrganizationInvitationUsecase, tokenService)
	// userOrganizationMembershipController := controllers.NewUserOrganizationMembershipController(userOrganizationMembershipUsecase, tokenService)

	r := router.SetUpRouter()

	user := r.Group("/user")
	{
		user.POST("/signup", userController.SignUp)
		user.POST("/signin", userController.SignIn)
		user.GET("/verify", userController.VerifyToken)
		user.GET("/organization", userController.GetJoinedOrganizations) // フロントでフェッチをまとめるために、招待を受けている組織もまとめて取得したい
		user.GET("/organization/invite", userController.GetRecievedInvitations)
	}

	organization := r.Group("/organization")
	{
		// organization
		organization.Use(middleware.JwtAuthMiddleware(tokenService))
		organization.POST("", organizationController.CreateOrganization)

		// invitation
		organization.POST("/invite", userOrganizationInvitationController.InviteUserToOrganization)
		organization.POST("/invite/cancel", userOrganizationInvitationController.CancelInvite)
		organization.GET("/:id/invite", userOrganizationInvitationController.GetSendInvitations)
		organization.POST("/invite/accept", userOrganizationInvitationController.AcceptInvite)
		organization.POST("/invite/reject", userOrganizationInvitationController.RejectInvite)
	}

	r.Run(":8080")
}