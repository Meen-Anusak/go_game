package core

import "github.com/google/uuid"

type UserService interface {
	GetAllUser([]User, error)
	CreateNewUser(user User) error
	GetUserById(id uuid.UUID) (User error)
}
