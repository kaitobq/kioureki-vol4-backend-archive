package usecases

import (
	"backend/domain/entities"
	"backend/domain/repository"
	"backend/usecases/response"
)

type OrganizationUsecase struct {
	OrganizationRepository               repository.OrganizationRepository
	UserOrganizationMembershipRepository repository.UserOrganizationMembershipRepository
	UserRepository                       repository.UserRepository
}

func NewOrganizationUsecase(organizationRepo repository.OrganizationRepository, userOrganizationMembershipRepo repository.UserOrganizationMembershipRepository, userRepo repository.UserRepository) *OrganizationUsecase {
	return &OrganizationUsecase{
		OrganizationRepository:               organizationRepo,
		UserOrganizationMembershipRepository: userOrganizationMembershipRepo,
		UserRepository:                       userRepo,
	}
}

func (ou * OrganizationUsecase) CreateOrganization(name string, userID uint) (*response.OrganizationResponse, error) {
	organization := &entities.Organization{Name: name}
	createdOrganization, err := ou.OrganizationRepository.Save(organization)
	if err != nil {
		return nil, err
	}

	err = ou.UserOrganizationMembershipRepository.AddToMemberships(organization.ID, userID)
	if err != nil {
		return nil, err
	}

	return response.NewOrganizationResponse(createdOrganization), nil
}


