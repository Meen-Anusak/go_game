package services

import (
	"go_game/crm/core/domain"
	"go_game/crm/core/repository"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	CreateNewCustomer(customer domain.Customer) (*domain.Customer, error)
}

type CustomerServiceImpl struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) *CustomerServiceImpl {
	return &CustomerServiceImpl{repo: repo}
}
func (s *CustomerServiceImpl) GetAllCustomer() ([]domain.Customer, error) {
	customers, err := s.repo.GetAllCustomer()
	if err != nil {
		return nil, err
	}
	var customerResponse []domain.Customer

	for _, customer := range customers {
		response := domain.Customer{
			CustomerID:   customer.CustomerID,
			CustomerName: customer.CustomerName,
			FistName:     customer.FistName,
			LastName:     customer.LastName,
			Email:        customer.Email,
			PhoneNumber:  customer.PhoneNumber,
			PlatFrom:     customer.PlatFrom,
			ActivityAt:   customer.ActivityAt,
		}
		customerResponse = append(customerResponse, response)
	}

	return customerResponse, nil
}

func (s *CustomerServiceImpl) CreateNewCustomer(customer domain.Customer) (*domain.Customer, error) {
	newCustomer, err := s.repo.CreateNewCustomer(customer)
	if err != nil {
		return nil, err
	}
	return newCustomer, nil
}
