package repository

import (
	"github.com/google/uuid"
	"go_game/game/core/domain"
)

type PlayerRepository interface {
	GetAllPlayer() ([]domain.Player, error)
	GetPlayerById(id uuid.UUID) (*domain.Player, error)
	CreateNewPlayer(player domain.Player) error
}
