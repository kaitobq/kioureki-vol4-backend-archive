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

func (ou *OrganizationUsecase) BeforeInvite(organizationID uint, email string, userID uint) error {
	sender, err := ou.UserRepository.FindByID(userID)
	if err != nil {
		return err
	}

	senderJoinedOrganization, err := ou.OrganizationRepository.CheckUserInOrganization(organizationID, sender.Email)
	if err != nil {
		return err
	}
	if !senderJoinedOrganization {
		return fmt.Errorf("%w(organizationID: %d)", errors.ErrUserNotInOrganization, organizationID)
	}

	userAlreadyInOrganization, err := ou.OrganizationRepository.CheckUserInOrganization(organizationID, email)
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

func (ou *OrganizationUsecase) GetInvitationsByUserID(userID uint) ([]repository.GetInvitedOrganizationsByUserIDOutput, error) {
	organizations, err := ou.OrganizationRepository.GetInvitedOrganizationsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}

func (ou *OrganizationUsecase) AcceptInvite(invitationID, userID uint) error {
	invitation, err := ou.OrganizationRepository.GetInvitationByID(invitationID)
	if err != nil {
		return err
	}
	if invitation.UserID != userID {
		return fmt.Errorf("%w", errors.ErrNoPermission)
	}

	err = ou.OrganizationRepository.AddToMemberships(invitation.ID, invitation.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (ou *OrganizationUsecase) RejectInvite(invitationID, userID uint) error {
	invitation, err := ou.OrganizationRepository.GetInvitationByID(invitationID)
	if err != nil {
		return err
	}
	if invitation.UserID != userID {
		return fmt.Errorf("%w", errors.ErrNoPermission)
	}

	err = ou.OrganizationRepository.DeleteInvitation(invitation.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ou *OrganizationUsecase) CancelInvite(invitationID, userID uint) error {
	invitation, err := ou.OrganizationRepository.GetInvitationByID(invitationID)
	if err != nil {
		return err
	}

	cancelerOrganizations, err := ou.OrganizationRepository.GetUserJoinedOrganization(userID)
	if err != nil {
		return err
	}

	count := 0
	for _, organization := range cancelerOrganizations {
		if organization.ID == invitation.OrganizationID {
			count++
		}
	}
	if count == 0 {
		return fmt.Errorf("%w", errors.ErrNoPermission)
	}

	err = ou.OrganizationRepository.DeleteInvitation(invitation.ID)
	if err != nil {
		return err
	}

	return nil
}
