package service

import (
	"go_game/auth/core/domain"
	"go_game/auth/core/repository"
)

type LoginService interface {
	Login(login domain.Login) error
}

type LoginImpl struct {
	repo repository.LoginRepository
}

func (s *LoginImpl) Login(login domain.Login) error {
	panic("login")
}
