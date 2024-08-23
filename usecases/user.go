package usecases

import (
	"backend/domain/entities"
	"backend/domain/errors"
	"backend/domain/repository"
	"backend/domain/service"
	"fmt"
)

type UserUsecase struct {
	UserRepository                       repository.UserRepository
	TokenService                         service.TokenService
	OrganizationRepository               repository.OrganizationRepository
	UserOrganizationInvitationRepository repository.UserOrganizationInvitationRepository
}

func NewUserUsecase(userRepo repository.UserRepository, tokenService *service.TokenService, organizationRepo repository.OrganizationRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepo,
		TokenService:   *tokenService,
		OrganizationRepository: organizationRepo,
	}
}

func (u *UserUsecase) CreateUser(name, email, password string) (*entities.User, string, error) {
	// メールが既に使用されていないか確認
	emailInUse, err := u.UserRepository.CheckEmailAlreadyInUse(email)
	if err != nil {
		return nil, "", err
	}
	if emailInUse {
		return nil, "", fmt.Errorf("%w(%s)", errors.ErrEmailAlreadyInUse, email)
	}

	// パスワードをハッシュ化
	hashedPassword, err := u.UserRepository.HashPassword(password)
	if err != nil {
		return nil, "", err
	}

	// 新しいユーザーを作成して保存
	user := &entities.User{Name: name, Email: email, Password: hashedPassword}
	err = u.UserRepository.Save(user)
	if err != nil {
		return nil, "", err
	}

	// トークンを生成
	token, err := u.TokenService.GenerateTokenFromID(uint(user.ID))
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (u *UserUsecase) Authenticate(email, password string) (*entities.User, string, error) {
	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, "", err
	}

	// パスワードの照合
	match, err := u.UserRepository.CheckPassword(user.Password, password)
	if err != nil {
		return nil, "", err
	}
	if !match {
		return nil, "", fmt.Errorf("%w", errors.ErrInvalidPassword)
	}

	// トークンを生成
	token, err := u.TokenService.GenerateTokenFromID(uint(user.ID))
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (u *UserUsecase) GetUserJoinedOrganization(userID uint) ([]entities.Organization, error) {
	organizations, err := u.OrganizationRepository.GetUserJoinedOrganization(userID)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}

func (u *UserUsecase) GetRecievedInvitationsByUserID(userID uint) ([]repository.GetRecievedInvitationsByUserIDOutput, error) {
	organizations, err := u.UserOrganizationInvitationRepository.GetRecievedInvitationsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}
