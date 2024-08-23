package repository

import (
	"backend/domain/entities"
	"database/sql"
)

type OrganizationRepository interface {
	Save(organization *entities.Organization) (*entities.Organization, error)
	CheckUserInOrganization(organizationID uint, email string) (bool, error)
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

func (r *MySQLOrganizationRepository) Save(organization *entities.Organization) (*entities.Organization, error) {
	result, err := r.DB.Exec("INSERT INTO organizations (name) VALUES (?)", organization.Name)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	organization.ID = uint(id)

	return organization, nil
}

func (r *MySQLOrganizationRepository) CheckUserInOrganization(organizationID uint, email string) (bool, error) {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM user_organization_memberships uom JOIN users u ON uom.user_id = u.id WHERE u.email = ? AND uom.organization_id = ?", email, organizationID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
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
