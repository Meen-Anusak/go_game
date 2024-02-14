package core

import "github.com/google/uuid"

type PlayerRepository interface {
	GetAllPlayer() ([]Player, error)
	GetPlayerById(id uuid.UUID) (*Player, error)
	CreateNewPlayer(player Player) error
}
