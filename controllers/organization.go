package controllers

import (
	"backend/domain/service"
	"backend/usecases"
	"backend/usecases/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	OrganizationUsecase *usecases.OrganizationUsecase
	TokenService        *service.TokenService
	UserUsecase         *usecases.UserUsecase
}

func NewOrganizationController(ou *usecases.OrganizationUsecase) *OrganizationController {
	return &OrganizationController{
		OrganizationUsecase: ou,
	}
}

func (oc *OrganizationController) CreateOrganization(c *gin.Context) {
	req, err := request.NewCreateOrganizationRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organization, err := oc.OrganizationUsecase.CreateOrganization(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userID, err := oc.TokenService.ExtractIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = oc.OrganizationUsecase.AddToMemberships(organization.ID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"organization": organization})
}

