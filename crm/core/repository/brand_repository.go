package repository

import "go_game/crm/core/domain"

type BrandRepository interface {
	GetAllBrand() ([]domain.Brand, error)
	CreateNewBrand(brand domain.Brand) (*domain.Brand, error)
}
