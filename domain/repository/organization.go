package repository

import (
	"backend/domain/entities"
	"database/sql"
)

type OrganizationRepository interface {
	Save(organization *entities.Organization) error
	AddToMemberships(organizationID, userID uint) error
	AddToInvitations(organizationID, userID uint) error
	CheckUserInOrganization(organizationID uint, email string) (bool, error)
	CheckAlreadySentInvite(organizationID uint, email string) (bool, error)
	GetRecievedInvitationsByUserID(userID uint) ([]GetRecievedInvitationsByUserIDOutput, error)
	GetSendInvitationsByOrganizationID(userID uint) ([]entities.UserOrganizationInvitation, error)
	GetInvitationByID(invitationID uint) (*entities.UserOrganizationInvitation, error)
	DeleteInvitation(invitationID uint) error
	GetUserJoinedOrganization(userID uint) ([]entities.Organization, error)
	FindByID(organizationID uint) (*entities.Organization, error)
}

type MySQLOrganizationRepository struct {
	DB *sql.DB
}

func NewMySQLOrganizationRepository(db *sql.DB) *MySQLOrganizationRepository {
	return &MySQLOrganizationRepository{
		DB: db,
	}
}

func (r *MySQLOrganizationRepository) Save(organization *entities.Organization) error {
	result, err := r.DB.Exec("INSERT INTO organizations (name) VALUES (?)", organization.Name)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	organization.ID = uint(id)

	return nil
}

func (r *MySQLOrganizationRepository) AddToMemberships(organizationID, userID uint) error {
	_, err := r.DB.Exec("INSERT INTO user_organization_memberships (organization_id, user_id) VALUES (?, ?)", organizationID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLOrganizationRepository) AddToInvitations(organizationID, userID uint) error {
	_, err := r.DB.Exec("INSERT INTO user_organization_invitations (organization_id, user_id) VALUES (?, ?)", organizationID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLOrganizationRepository) CheckUserInOrganization(organizationID uint, email string) (bool, error) {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM user_organization_memberships uom JOIN users u ON uom.user_id = u.id WHERE u.email = ? AND uom.organization_id = ?", email, organizationID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *MySQLOrganizationRepository) CheckAlreadySentInvite(organizationID uint, email string) (bool, error) {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM user_organization_invitations uoi JOIN users u ON uoi.user_id = u.id WHERE u.email = ? AND uoi.organization_id = ?", email, organizationID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

type GetRecievedInvitationsByUserIDOutput struct {
	OrganizationID uint `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	InvitationID uint `json:"invitation_id"`
}

func (r *MySQLOrganizationRepository) GetRecievedInvitationsByUserID(userID uint) ([]GetRecievedInvitationsByUserIDOutput, error) {
	rows, err := r.DB.Query("SELECT o.id, o.name, uoi.id FROM organizations o JOIN user_organization_invitations uoi ON o.id = uoi.organization_id WHERE uoi.user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var outputs []GetRecievedInvitationsByUserIDOutput
	for rows.Next() {
		var output GetRecievedInvitationsByUserIDOutput
		err := rows.Scan(&output.OrganizationID, &output.OrganizationName, &output.InvitationID)
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, output)
	}

	return outputs, nil
}

func (r *MySQLOrganizationRepository) GetSendInvitationsByOrganizationID(organizationID uint) ([]entities.UserOrganizationInvitation, error) {
	rows, err := r.DB.Query("SELECT id, user_id, organization_id FROM user_organization_invitations WHERE organization_id = ?", organizationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invitations []entities.UserOrganizationInvitation
	for rows.Next() {
		var invitation entities.UserOrganizationInvitation
		err := rows.Scan(&invitation.ID, &invitation.UserID, &invitation.OrganizationID)
		if err != nil {
			return nil, err
		}
		invitations = append(invitations, invitation)
	}

	return invitations, nil
}

func (r *MySQLOrganizationRepository) GetInvitationByID(invitationID uint) (*entities.UserOrganizationInvitation, error) {
	var invitation entities.UserOrganizationInvitation
	err := r.DB.QueryRow("SELECT id, user_id, organization_id FROM user_organization_invitations WHERE id = ?", invitationID).Scan(&invitation.ID, &invitation.UserID, &invitation.OrganizationID)
	if err != nil {
		return nil, err
	}

	return &invitation, nil
}

func (r *MySQLOrganizationRepository) DeleteInvitation(invitationID uint) error {
	row := r.DB.QueryRow("DELETE FROM user_organization_invitations WHERE id = ?", invitationID)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}

func (r *MySQLOrganizationRepository) GetUserJoinedOrganization(userID uint) ([]entities.Organization, error) {
	rows, err := r.DB.Query("SELECT o.id, o.name FROM organizations o JOIN user_organization_memberships uom ON o.id = uom.organization_id WHERE uom.user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var organizations []entities.Organization
	for rows.Next() {
		var organization entities.Organization
		err := rows.Scan(&organization.ID, &organization.Name)
		if err != nil {
			return nil, err
		}
		organizations = append(organizations, organization)
	}

	return organizations, nil
}

func (r *MySQLOrganizationRepository) FindByID(organizationID uint) (*entities.Organization, error) {
	var organization entities.Organization
	err := r.DB.QueryRow("SELECT id, name FROM organizations WHERE id = ?", organizationID).Scan(&organization.ID, &organization.Name)
	if err != nil {
		return nil, err
	}

	return &organization, nil
}
