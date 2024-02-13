package core

import "github.com/google/uuid"

type PlayerRepository interface {
	GetAll([]Player) error
	GetById(id uuid.UUID) (*Player, error)
	Create(player Player) error
	Update(id uuid.UUID) (*Player, error)
	Delete(id uuid.UUID) (*Player, error)
}
