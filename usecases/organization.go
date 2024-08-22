package usecases

import (
	"backend/domain/entities"
	"backend/domain/errors"
	"backend/domain/repository"
	"fmt"
)

type OrganizationUsecase struct {
	OrganizationRepository repository.OrganizationRepository
	UserRepository         repository.UserRepository
}

func NewOrganizationUsecase(organizationRepo repository.OrganizationRepository, userRepo repository.UserRepository) *OrganizationUsecase {
	return &OrganizationUsecase{
		OrganizationRepository: organizationRepo,
		UserRepository: userRepo,
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
	err := ou.OrganizationRepository.AddToMemberships(organizationID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (ou *OrganizationUsecase) BeforeInvite(organizationID uint, email string) error {
	userAlreadyInOrganization, err := ou.OrganizationRepository.CheckUserAlreadyInOrganization(organizationID, email)
	if err != nil {
		return err
	}
	if userAlreadyInOrganization {
		return fmt.Errorf("%w", errors.ErrUserAlreadyInOrganization)
	}

	alreadySentInvite, err := ou.OrganizationRepository.CheckAlreadySentInvite(organizationID, email)
	if err != nil {
		return err
	}
	if alreadySentInvite {
		return fmt.Errorf("%w", errors.ErrInviteAlreadySent)
	}

	return nil
}

func (ou *OrganizationUsecase) InviteUserToOrganization(organizationID uint, email string) error {
	user, err := ou.UserRepository.FindByEmail(email)
	if err != nil {
		return err
	}

	err = ou.OrganizationRepository.AddToInvitations(organizationID, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ou *OrganizationUsecase) GetInvitationsByUserID(userID uint) ([]entities.Organization, error) {
	organizations, err := ou.OrganizationRepository.GetInvitationsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}
