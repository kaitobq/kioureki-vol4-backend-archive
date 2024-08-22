package entities

type UserOrganizationInvitation struct {
	ID             int `json:"id"`
	UserID         int `json:"user_id"`
	OrganizationID int `json:"organization_id"`
}