package repository

import "go_game/auth/core/domain"

type LoginRepository interface {
	Login(login domain.Login) error
}
