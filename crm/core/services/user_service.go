package services

import (
	"go_game/crm/core/domain"
	"go_game/crm/core/repository"
)

type UserService interface {
	CreateNewUser(user domain.User) error
}
type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) CreateNewUser(user domain.User) error {
	if err := s.repo.CreateNewUser(user); err != nil {
		return err
	}
	return nil
}
