package repository

import "database/sql"

type OrganizationRepository interface {
}

type MySQLOrganizationRepository struct {
	DB *sql.DB
}

func NewMySQLOrganizationRepository(db *sql.DB) *MySQLOrganizationRepository {
	return &MySQLOrganizationRepository{
		DB: db,
	}
}
