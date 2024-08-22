package controllers

import "backend/usecases"

type OrganizationController struct {
	OrganizationUsecase *usecases.OrganizationUsecase
}

func NewOrganizationController(ou *usecases.OrganizationUsecase) *OrganizationController {
	return &OrganizationController{
		OrganizationUsecase: ou,
	}
}
