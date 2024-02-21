package repository

import (
	"go_game/game/core/domain"
)

type PlayerRepository interface {
	GetAllPlayer() ([]domain.Player, error)
	GetPlayerById(id string) (*domain.Player, error)
	CreateNewPlayer(player domain.Player) error
	UpdatePlayer(player domain.Player) (*domain.Player, error)
	DeletePlayer(id string) error
}
