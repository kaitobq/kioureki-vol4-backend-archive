package entities

type UserOrganizationInvitation struct {
	ID             uint `json:"id"`
	UserID         uint `json:"user_id"`
	OrganizationID uint `json:"organization_id"`
}