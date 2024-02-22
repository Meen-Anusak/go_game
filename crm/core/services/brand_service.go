package services

import (
	"go_game/crm/core/domain"
	"go_game/crm/core/repository"
)

type BrandService interface {
	GetAllBrand() ([]domain.Brand, error)
	CreateNewBrand(brand domain.Brand) (*domain.Brand, error)
}
type BrandServiceImpl struct {
	repo repository.BrandRepository
}

func NewBrandService(repo repository.BrandRepository) *BrandServiceImpl {
	return &BrandServiceImpl{repo: repo}
}

func (s *BrandServiceImpl) GetAllBrand() ([]domain.Brand, error) {
	brands, err := s.repo.GetAllBrand()
	if err != nil {
		return nil, err
	}

	var brandResponse []domain.Brand

	for _, brand := range brands {
		response := domain.Brand{
			BrandID:   brand.BrandID,
			BrandName: brand.BrandName,
			CreatedAt: brand.CreatedAt,
			UpdatedAt: brand.UpdatedAt,
		}
		brandResponse = append(brandResponse, response)
	}
	return brandResponse, nil
}

func (s *BrandServiceImpl) CreateNewBrand(brand domain.Brand) (*domain.Brand, error) {
	newBrand, err := s.repo.CreateNewBrand(brand)
	if err != nil {
		return nil, err
	}
	return newBrand, nil
}
