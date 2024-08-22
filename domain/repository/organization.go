package repository

import (
	"backend/domain/entities"
	"database/sql"
)

type OrganizationRepository interface {
	Save(organization *entities.Organization) error
	AddToMemberships(organizationID, userID uint) error
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
