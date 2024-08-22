package repository

import (
	"backend/domain/entities"
	"database/sql"
)

type OrganizationRepository interface {
	Save(organization *entities.Organization) error
	AddToMemberships(organizationID, userID uint) error
	AddToInvitations(organizationID, userID uint) error
	CheckUserAlreadyInOrganization(organizationID uint, email string) (bool, error)
	CheckAlreadySentInvite(organizationID uint, email string) (bool, error)
	GetInvitationsByUserID(userID uint) ([]entities.Organization, error)
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

func (r *MySQLOrganizationRepository) CheckUserAlreadyInOrganization(organizationID uint, email string) (bool, error) {
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

func (r *MySQLOrganizationRepository) GetInvitationsByUserID(userID uint) ([]entities.Organization, error) {
	rows, err := r.DB.Query("SELECT o.id, o.name FROM organizations o JOIN user_organization_invitations uoi ON o.id = uoi.organization_id WHERE uoi.user_id = ?", userID)
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
