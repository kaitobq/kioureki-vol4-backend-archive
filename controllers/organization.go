package controllers

import (
	"backend/domain/service"
	"backend/usecases"
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

type createOrganizationInput struct {
	Name string `json:"name" binding:"required"`
}

func (oc *OrganizationController) CreateOrganization(c *gin.Context) {
	var input createOrganizationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organization, err := oc.OrganizationUsecase.CreateOrganization(input.Name)
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

type InviteUserToOrganizationInput struct {
	OrganizationId uint   `json:"organization_id" binding:"required"`
	Email 		   string `json:"email" binding:"required"`
}

func (oc *OrganizationController) InviteUserToOrganization(c *gin.Context) {
	var input InviteUserToOrganizationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := oc.OrganizationUsecase.BeforeInvite(input.OrganizationId, input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = oc.OrganizationUsecase.InviteUserToOrganization(input.OrganizationId, input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invitation sent"})
}

func (oc *OrganizationController) GetRecievedInvitations(c *gin.Context) {
	userID, err := oc.TokenService.ExtractIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	organizations, err := oc.OrganizationUsecase.GetInvitationsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"organizations": organizations})
}
