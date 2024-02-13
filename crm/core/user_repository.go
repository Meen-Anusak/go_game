package core

import "github.com/google/uuid"

type UserRepository interface {
	GetAll([]User) error
	GetById(id uuid.UUID) (*User, error)
	Create(player User) error
	Update(id uuid.UUID) (*User, error)
	Delete(id uuid.UUID) (*User, error)
}
