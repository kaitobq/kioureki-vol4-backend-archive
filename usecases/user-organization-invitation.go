package usecases

import (
	"backend/domain/entities"
	"backend/domain/errors"
	"backend/domain/repository"
	"backend/usecases/response"
	"fmt"
)

type UserOrganizationInvitationUsecase struct {
	UserOrganizationInvitationRepository repository.UserOrganizationInvitationRepository
	UserOrganizationMembershipRepository repository.UserOrganizationMembershipRepository
	OrganizationRepository               repository.OrganizationRepository
	UserRepository 					     repository.UserRepository
}

func NewUserOrganizationInvitationUsecase(userOrganizationInvitationRepo repository.UserOrganizationInvitationRepository, userOrganizationMembershipRepo repository.UserOrganizationMembershipRepository, organizationRepo repository.OrganizationRepository, userRepo repository.UserRepository) *UserOrganizationInvitationUsecase {
	return &UserOrganizationInvitationUsecase{
		UserOrganizationInvitationRepository: userOrganizationInvitationRepo,
		UserOrganizationMembershipRepository: userOrganizationMembershipRepo,
		OrganizationRepository:               organizationRepo,
		UserRepository:                       userRepo,
	}
}


func (uoiu *UserOrganizationInvitationUsecase) BeforeInvite(organizationID uint, email string, userID uint) error {
	sender, err := uoiu.UserRepository.FindByID(userID)
	if err != nil {
		return err
	}

	senderJoinedOrganization, err := uoiu.OrganizationRepository.CheckUserInOrganization(organizationID, sender.Email)
	if err != nil {
		return err
	}
	if !senderJoinedOrganization {
		return fmt.Errorf("%w(organizationID: %d)", errors.ErrUserNotInOrganization, organizationID)
	}

	userAlreadyInOrganization, err := uoiu.OrganizationRepository.CheckUserInOrganization(organizationID, email)
	if err != nil {
		return err
	}
	if userAlreadyInOrganization {
		return fmt.Errorf("%w", errors.ErrUserAlreadyInOrganization)
	}

	alreadySentInvite, err := uoiu.UserOrganizationInvitationRepository.CheckAlreadySentInvite(organizationID, email)
	if err != nil {
		return err
	}
	if alreadySentInvite {
		return fmt.Errorf("%w", errors.ErrInviteAlreadySent)
	}

	return nil
}

func (uoiu *UserOrganizationInvitationUsecase) InviteUserToOrganization(organizationID uint, email string) error {
	user, err := uoiu.UserRepository.FindByEmail(email)
	if err != nil {
		return err
	}

	err = uoiu.UserOrganizationInvitationRepository.AddToInvitations(organizationID, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (uoiu *UserOrganizationInvitationUsecase) GetSendInvitationsByOrganizationID(organizationID uint) (*response.UserOrganizationInvitationsResponse, error) {
	invitations, err := uoiu.UserOrganizationInvitationRepository.GetSendInvitationsByOrganizationID(organizationID)
	if err != nil {
		return nil, err
	}

	organization, err := uoiu.OrganizationRepository.FindByID(organizationID)
	if err != nil {
		return nil, err
	}

	var users []entities.User
	for _, invitation := range invitations {
		user, err := uoiu.UserRepository.FindByID(invitation.UserID)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}

	return response.NewUserOrganizationInvitationsResponse(*organization, users, invitations), nil
}

func (uoiu *UserOrganizationInvitationUsecase) AcceptInvite(invitationID, userID uint) (*response.OrganizationResponse, error) {
	invitation, err := uoiu.UserOrganizationInvitationRepository.GetInvitationByID(invitationID)
	if err != nil {
		return nil, err
	}
	if invitation.UserID != userID {
		return nil, fmt.Errorf("%w", errors.ErrNoPermission)
	}

	err = uoiu.UserOrganizationMembershipRepository.AddToMemberships(invitation.ID, invitation.UserID)
	if err != nil {
		return nil, err
	}

	organization, err := uoiu.OrganizationRepository.FindByID(invitation.OrganizationID)
	if err != nil {
		return nil, err
	}

	return response.NewOrganizationResponse(organization), nil
}


func (uoiu *UserOrganizationInvitationUsecase) CancelInvite(invitationID, userID uint) error {
	invitation, err := uoiu.UserOrganizationInvitationRepository.GetInvitationByID(invitationID)
	if err != nil {
		return err
	}

	cancelerOrganizations, err := uoiu.OrganizationRepository.GetUserJoinedOrganization(userID)
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

	err = uoiu.UserOrganizationInvitationRepository.DeleteInvitation(invitation.ID)
	if err != nil {
		return err
	}

	return nil
}

func (uoiu *UserOrganizationInvitationUsecase) RejectInvite(invitationID, userID uint) (*response.OrganizationResponse, error) {
	invitation, err := uoiu.UserOrganizationInvitationRepository.GetInvitationByID(invitationID)
	if err != nil {
		return nil, err
	}
	if invitation.UserID != userID {
		return nil, fmt.Errorf("%w", errors.ErrNoPermission)
	}

	err = uoiu.UserOrganizationInvitationRepository.DeleteInvitation(invitation.ID)
	if err != nil {
		return nil, err
	}

	rejectedOrganization, err := uoiu.OrganizationRepository.FindByID(invitation.OrganizationID)
	if err != nil {
		return nil, err
	}

	return response.NewOrganizationResponse(rejectedOrganization), nil
}
