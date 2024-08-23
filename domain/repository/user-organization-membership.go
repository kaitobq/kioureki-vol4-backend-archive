package repository

import "database/sql"

type UserOrganizationMembershipRepository interface {
	AddToMemberships(organizationID, userID uint) error
}

type MySQLUserOrganizationMembershipRepository struct {
	DB *sql.DB
}

func NewMySQLUserOrganizationMembershipRepository(db *sql.DB) *MySQLUserOrganizationMembershipRepository {
	return &MySQLUserOrganizationMembershipRepository{
		DB: db,
	}
}

func (r *MySQLUserOrganizationMembershipRepository) AddToMemberships(organizationID, userID uint) error {
	_, err := r.DB.Exec("INSERT INTO user_organization_memberships (organization_id, user_id) VALUES (?, ?)", organizationID, userID)
	if err != nil {
		return err
	}
	return nil
}