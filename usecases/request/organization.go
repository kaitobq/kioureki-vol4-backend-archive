package request

import "github.com/gin-gonic/gin"

type CreateOrganizationRequest struct {
	Name string `json:"name" binding:"required"`
}

func NewCreateOrganizationRequest(c *gin.Context) (*CreateOrganizationRequest, error) {
	var request CreateOrganizationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	return &CreateOrganizationRequest{
		Name: request.Name,
	}, nil
}

type InviteUserToOrganizationRequest struct {
	OrganizationId uint   `json:"organization_id" binding:"required"`
	Email 		   string `json:"email" binding:"required"`
}

func NewInviteUserToOrganizationRequest(c *gin.Context) (*InviteUserToOrganizationRequest, error) {
	var request InviteUserToOrganizationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	return &InviteUserToOrganizationRequest{
		OrganizationId: request.OrganizationId,
		Email: request.Email,
	}, nil
}
