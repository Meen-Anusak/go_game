package repository

import "go_game/crm/core/domain"

type CustomerRepository interface {
	GetAllCustomer() ([]domain.Customer, error)
	CreateNewCustomer(customer domain.Customer) (*domain.Customer, error)
}
