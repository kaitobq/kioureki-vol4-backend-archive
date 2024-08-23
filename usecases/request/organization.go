package request

import "github.com/gin-gonic/gin"

type CreateOrganizationRequest struct {
	Name string `json:"name" binding:"required"`
}

func NewCreateOrganizationRequest(c *gin.Context) (*CreateOrganizationRequest, error) {
	var req CreateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &CreateOrganizationRequest{
		Name: req.Name,
	}, nil
}

type InviteUserToOrganizationRequest struct {
	OrganizationId uint   `json:"organization_id" binding:"required"`
	Email 		   string `json:"email" binding:"required"`
}

func NewInviteUserToOrganizationRequest(c *gin.Context) (*InviteUserToOrganizationRequest, error) {
	var req InviteUserToOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
