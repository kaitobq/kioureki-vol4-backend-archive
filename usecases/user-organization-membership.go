package usecases

import "backend/domain/repository"

type UserOrganizationMembershipUsecase struct {
	UserOrganizationMembershipRepository repository.UserOrganizationMembershipRepository
}

func NewUserOrganizationMembershipUsecase(userOrganizationMembershipRepo repository.UserOrganizationMembershipRepository) *UserOrganizationMembershipUsecase {
	return &UserOrganizationMembershipUsecase{
		UserOrganizationMembershipRepository: userOrganizationMembershipRepo,
	}
}
