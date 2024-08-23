package usecases

import (
	"backend/domain/entities"
	"backend/domain/repository"
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

func (ou * OrganizationUsecase) CreateOrganization(name string) (*entities.Organization, error) {
	organization := &entities.Organization{Name: name}
	err := ou.OrganizationRepository.Save(organization)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (ou *OrganizationUsecase) AddToMemberships(organizationID, userID uint) error {
	err := ou.UserOrganizationMembershipRepository.AddToMemberships(organizationID, userID)
	if err != nil {
		return err
	}

	return nil
}

