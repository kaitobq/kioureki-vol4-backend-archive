package usecases

import (
	"backend/domain/entities"
	"backend/domain/errors"
	"backend/domain/repository"
	"fmt"
)

type UserUsecase struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepo,
	}
}

func (u *UserUsecase) CreateUser(name, email, password string) (*entities.User, error) {
	// メールが既に使用されていないか確認
	emailInUse, err := u.UserRepository.CheckEmailAlreadyInUse(email)
	if err != nil {
		return nil, err
	}
	if emailInUse {
		return nil, fmt.Errorf("%w(%s)", errors.ErrEmailAlreadyInUse, email)
	}

	// パスワードをハッシュ化
	hashedPassword, err := u.UserRepository.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// 新しいユーザーを作成して保存
	user := &entities.User{Name: name, Email: email, Password: hashedPassword}
	err = u.UserRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecase) Authenticate(email, password string) (*entities.User, error) {
	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	// パスワードの照合
	match, err := u.UserRepository.CheckPassword(user.Password, password)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, fmt.Errorf("%w", errors.ErrInvalidPassword)
	}

	return user, nil
}