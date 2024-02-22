package handler

import (
	"go_game/crm/core/domain"
	"go_game/crm/core/repository"
	"gorm.io/gorm"
)

type GormCustomerRepository struct {
	db *gorm.DB
}

func NewGormCustomerRepository(db *gorm.DB) repository.CustomerRepository {
	return &GormCustomerRepository{db: db}
}

func (r *GormCustomerRepository) GetAllCustomer() ([]domain.Customer, error) {
	var customer []domain.Customer
	if result := r.db.Find(&customer); result.Error != nil {
		return nil, result.Error
	}
	return customer, nil
}

func (r *GormCustomerRepository) CreateNewCustomer(customer domain.Customer) (*domain.Customer, error) {
	result := r.db.Create(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}
