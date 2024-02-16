package repository

import "go_game/crm/core/domain"

type UserRepository interface {
	GetAllUser() ([]domain.User, error)
	GetUserById(id string) (*domain.User, error)
	CreateNewUser(user domain.User) error
}
