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

func (uoic *UserOrganizationInvitationController) AcceptInvite(c *gin.Context) {
	req, err := request.NewInviteRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uoic.TokenService.ExtractIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = uoic.UserOrganizationInvitationUsecase.AcceptInvite(req.InvitationID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invitation accepted"})
}

func (uoic *UserOrganizationInvitationController) RejectInvite(c *gin.Context) {
	req, err := request.NewInviteRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uoic.TokenService.ExtractIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = uoic.UserOrganizationInvitationUsecase.RejectInvite(req.InvitationID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invitation rejected"})
}

func (uoic *UserOrganizationInvitationController) CancelInvite(c *gin.Context) {
	req, err := request.NewInviteRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uoic.TokenService.ExtractIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = uoic.UserOrganizationInvitationUsecase.CancelInvite(req.InvitationID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invitation cancelled"})
}