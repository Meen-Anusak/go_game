package services

import (
	"go_game/game/core/domain"
	"go_game/game/core/repository"
)

type PlayerService interface {
	GetAllPlayer() ([]domain.Player, error)
	GetPlayerById(id string) (*domain.Player, error)
	CreateNewPlayer(player domain.Player) error
	UpdatePlayer(player domain.Player) (*domain.Player, error)
}

type PlayerServiceImpl struct {
	repo repository.PlayerRepository
}

func NewPlayerService(repo repository.PlayerRepository) *PlayerServiceImpl {
	return &PlayerServiceImpl{repo: repo}
}

func (s *PlayerServiceImpl) CreateNewPlayer(player domain.Player) error {
	if err := s.repo.CreateNewPlayer(player); err != nil {
		return err
	}
	return nil
}

func (s *PlayerServiceImpl) GetPlayerById(id string) (*domain.Player, error) {
	player, err := s.repo.GetPlayerById(id)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (s *PlayerServiceImpl) GetAllPlayer() ([]domain.Player, error) {
	player, err := s.repo.GetAllPlayer()
	if err != nil {
		return nil, err
	}

	var playerResponses []domain.Player

	for _, player := range player {
		response := domain.Player{
			PlayerID:   player.PlayerID,
			PlayerName: player.PlayerName,
			CreatedAt:  player.CreatedAt,
			UpdatedAt:  player.UpdatedAt,
		}
		playerResponses = append(playerResponses, response)
	}

	return playerResponses, nil
}

func (s *PlayerServiceImpl) UpdatePlayer(player domain.Player) (*domain.Player, error) {
	result, err := s.repo.UpdatePlayer(player)
	if err != nil {
		return nil, err
	}
	return result, nil
}
