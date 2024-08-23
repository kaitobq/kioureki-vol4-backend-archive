package response

type UserOrganizationInvitationResponse struct {
	OrganizationID   uint   `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	UserID		     uint   `json:"user_id"`
	UserName 	     string `json:"user_name"`
	InvitationID     uint   `json:"invitation_id"`
}

func NewUserOrganizationInvitationResponse(organizationID uint, organizationName string, userID uint, userName string, invitationID uint) *UserOrganizationInvitationResponse {
	return &UserOrganizationInvitationResponse{
		OrganizationID:   organizationID,
		OrganizationName: organizationName,
		UserID: 		  userID,
		UserName:  		  userName,
		InvitationID:     invitationID,
	}
}
