package repository

import "go_game/auth/core/domain"

type AuthRepository interface {
	Login(login domain.Login) (*domain.User, error)

	Register(register domain.User) error
}
