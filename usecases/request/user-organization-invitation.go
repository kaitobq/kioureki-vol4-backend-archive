package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetSendInvitationsRequest struct {
	OrganizationID uint `json:"organization_id"`
}

func NewGetSendInvitationsRequest(c *gin.Context) *GetSendInvitationsRequest {
	id := c.Param("id")
	
	organizationID, _ := strconv.ParseUint(id, 10, 32)
	
	return &GetSendInvitationsRequest{
		OrganizationID: uint(organizationID),
	}
}
