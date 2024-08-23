package response

import "backend/domain/entities"

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

type UserInvitation struct {
	UserID       uint   `json:"user_id"`
	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	InvitationID uint   `json:"invitation_id"`
}

type UserOrganizationInvitationsResponse struct {
	OrganizationID     uint           `json:"organization_id"`
	OrganizationName   string         `json:"organization_name"`
	Users            []UserInvitation `json:"users"`
}

func NewUserOrganizationInvitationsResponse(organization entities.Organization, users []entities.User, invitations []entities.UserOrganizationInvitation) *UserOrganizationInvitationsResponse {
	var userInvitations []UserInvitation

	// マップを作成して、ユーザーのIDとそのインビテーションを簡単に一致させる
	invitationMap := make(map[uint]entities.UserOrganizationInvitation)
	for _, invitation := range invitations {
		invitationMap[invitation.UserID] = invitation
	}

	// ユーザーと対応するインビテーションを一致させてリストに追加
	for _, user := range users {
		if invitation, ok := invitationMap[user.ID]; ok {
			userInvitations = append(userInvitations, UserInvitation{
				UserID:       user.ID,
				UserName:     user.Name,
				UserEmail:    user.Email,
				InvitationID: invitation.ID,
			})
		}
	}

	return &UserOrganizationInvitationsResponse{
		OrganizationID:   organization.ID,
		OrganizationName: organization.Name,
		Users:            userInvitations,
	}
}
