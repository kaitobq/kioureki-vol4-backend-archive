package response

import "backend/domain/entities"

type OrganizationResponse struct {
	ID          uint `json:"id"`
	Name        string `json:"name"`
}

func NewOrganizationResponse(organization *entities.Organization) *OrganizationResponse {
	return &OrganizationResponse{
		ID:   organization.ID,
		Name: organization.Name,
	}
}
