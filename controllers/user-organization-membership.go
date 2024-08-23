package controllers

import (
	"backend/domain/service"
	"backend/usecases"
)

type UserOrganizationMembershipController struct {
	UserOrganizationMembershipUsecase *usecases.UserOrganizationMembershipUsecase
	TokenService                      *service.TokenService
}

func NewUserOrganizationMembershipController(uomc *usecases.UserOrganizationMembershipUsecase, ts *service.TokenService) *UserOrganizationMembershipController {
	return &UserOrganizationMembershipController{
		UserOrganizationMembershipUsecase: uomc,
		TokenService:                      ts,
	}
}
