package controllers

import (
	"backend/domain/service"
	"backend/usecases"
	"backend/usecases/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserOrganizationInvitationController struct {
	UserOrganizationInvitationUsecase usecases.UserOrganizationInvitationUsecase
	TokenService                      *service.TokenService
}

func NewUserOrganizationInvitationController(uoic *usecases.UserOrganizationInvitationUsecase, ts *service.TokenService) *UserOrganizationInvitationController {
	return &UserOrganizationInvitationController{
		UserOrganizationInvitationUsecase: *uoic,
		TokenService:                      ts,
	}
}

func (uoic *UserOrganizationInvitationController) InviteUserToOrganization(c *gin.Context) {
	req, err := request.NewInviteUserToOrganizationRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uoic.TokenService.ExtractIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = uoic.UserOrganizationInvitationUsecase.BeforeInvite(req.OrganizationId, req.Email, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = uoic.UserOrganizationInvitationUsecase.InviteUserToOrganization(req.OrganizationId, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invitation sent"})
}

func (uoic *UserOrganizationInvitationController) GetSendInvitations(c *gin.Context) {
	req := request.NewGetSendInvitationsRequest(c)

	response, err := uoic.UserOrganizationInvitationUsecase.GetSendInvitationsByOrganizationID(req.OrganizationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"invitations": response})
}


type acceptInviteInput struct {
	InvitationID uint `json:"invitation_id" binding:"required"`
}

func (uoic *UserOrganizationInvitationController) AcceptInvite(c *gin.Context) {
	var input acceptInviteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uoic.TokenService.ExtractIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = uoic.UserOrganizationInvitationUsecase.AcceptInvite(input.InvitationID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invitation accepted"})
}

type rejectInviteInput struct {
	InvitationID uint `json:"invitation_id" binding:"required"`
}

func (uoic *UserOrganizationInvitationController) RejectInvite(c *gin.Context) {
	var input rejectInviteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errpr": err})
		return
	}

	userID, err := uoic.TokenService.ExtractIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = uoic.UserOrganizationInvitationUsecase.RejectInvite(input.InvitationID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invitation rejected"})
}

type cancelInviteInput struct {
	InvitationID uint `json:"invitation_id" binding:"required"`
}

func (uoic *UserOrganizationInvitationController) CancelInvite(c *gin.Context) {
	var input cancelInviteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uoic.TokenService.ExtractIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = uoic.UserOrganizationInvitationUsecase.CancelInvite(input.InvitationID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invitation cancelled"})
}