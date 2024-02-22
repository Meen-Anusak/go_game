package handler

import (
	"go_game/crm/core/domain"
	"go_game/crm/core/repository"
	"gorm.io/gorm"
)

type GormBrandRepository struct {
	db *gorm.DB
}

func NewGormBrandRepository(db *gorm.DB) repository.BrandRepository {
	return &GormBrandRepository{db: db}
}

func (r *GormBrandRepository) GetAllBrand() ([]domain.Brand, error) {
	var brand []domain.Brand
	if result := r.db.Find(&brand); result.Error != nil {
		return nil, result.Error
	}
	return brand, nil
}

func (r *GormBrandRepository) CreateNewBrand(brand domain.Brand) (*domain.Brand, error) {
	result := r.db.Create(&brand)
	if result.Error != nil {
		return nil, result.Error
	}
	return &brand, nil
}
