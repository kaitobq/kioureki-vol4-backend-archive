package response

import "backend/domain/entities"

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type AuthedUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}

func NewAuthedUserResponse(user entities.User, token string) *AuthedUserResponse {
	return &AuthedUserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Token: token,
	}
}

type UserJoinedOrganizationResponse struct {
	User             UserResponse `json:"user"`
	Organizatioins []OrganizationResponse `json:"organizations"`
}

func NewUserJoinedOrganizationResponse(user *entities.User, organizations []entities.Organization) *UserJoinedOrganizationResponse {
	orgs := make([]OrganizationResponse, len(organizations))
	for i, org := range organizations {
		orgs[i] = OrganizationResponse{
			ID:   org.ID,
			Name: org.Name,
		}
	}
	return &UserJoinedOrganizationResponse{
		User: UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
		Organizatioins: orgs,
	}
}
