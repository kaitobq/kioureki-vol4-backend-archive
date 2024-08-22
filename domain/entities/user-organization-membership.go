package entities

type UserOrganizationMembership struct {
	ID             int `json:"id"`
	UserID         int `json:"user_id"`
	OrganizationID int `json:"organization_id"`
}