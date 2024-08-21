// /domain/repository/user_repository.go
package repository

import (
	"backend/domain/entities"
	"backend/domain/errors"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
    FindByID(id uint) (*entities.User, error)
    FindByEmail(email string) (*entities.User, error)
    CheckEmailAlreadyInUse(email string) (bool, error)
    HashPassword(password string) (string, error)
    CheckPassword(password string) (bool, error)
    Save(user *entities.User) error
}

type MySQLUserRepository struct {
    DB *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
    return &MySQLUserRepository{
        DB: db,
    }
}

func (r *MySQLUserRepository) FindByID(id uint) (*entities.User, error) {
    var user entities.User
    err := r.DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *MySQLUserRepository) FindByEmail(email string) (*entities.User, error) {
    var user entities.User
    err := r.DB.QueryRow("SELECT id, name, email FROM users WHERE email = ?", email).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *MySQLUserRepository) CheckEmailAlreadyInUse(email string) (bool, error) {
    var count int
    err := r.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
    if err != nil {
        return false, err
    }
    return count > 0, nil
}

func (r *MySQLUserRepository) HashPassword(password string) (string, error) {
    if password == "" || len(password) == 0 {
        return "", fmt.Errorf("%w", errors.ErrPasswordEmpty)
    }

    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        return "", err
    }

    return string(bytes), nil
}

func (r *MySQLUserRepository) CheckPassword(password string) (bool, error) {
    if password == "" || len(password) == 0 {
        return false, fmt.Errorf("%w", errors.ErrPasswordEmpty)
    }

    err := bcrypt.CompareHashAndPassword([]byte(password), []byte(password))
    if err != nil {
        return false, err
    }

    return false, nil
}

func (r *MySQLUserRepository) Save(user *entities.User) error {
    _, err := r.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
    return err
}
