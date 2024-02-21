package services

import (
	"go_game/crm/core/domain"
	"go_game/crm/core/repository"
)

type UserService interface {
	CreateNewUser(user domain.User) error
	GatAllUser() ([]domain.User, error)
}
type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GatAllUser() ([]domain.User, error) {
	users, err := s.repo.GetAllUser()
	if err != nil {
		return nil, err
	}
	var userResponse []domain.User

	for _, user := range users {
		response := domain.User{
			UserID:      user.UserID,
			UserName:    user.UserName,
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			PhoneNumber: user.PhoneNumber,
			UpdatedAt:   user.UpdatedAt,
			CreatedAt:   user.CreatedAt,
		}
		userResponse = append(userResponse, response)
	}

	return userResponse, nil
}

func (s *UserServiceImpl) CreateNewUser(user domain.User) error {
	if err := s.repo.CreateNewUser(user); err != nil {
		return err
	}
	return nil
}
