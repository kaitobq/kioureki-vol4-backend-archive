package usecases

import "backend/domain/repository"

type OrganizationUsecase struct {
	OrganizationRepository repository.OrganizationRepository
}

func NewOrganizationUsecase(organizationRepo repository.OrganizationRepository) *OrganizationUsecase {
	return &OrganizationUsecase{
		OrganizationRepository: organizationRepo,
	}
}