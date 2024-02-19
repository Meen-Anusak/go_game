package core

import (
	"github.com/google/uuid"
	"go_game/game/core/domain"
	"go_game/game/core/repository"
)

type PlayerService interface {
	GetAllPlayer() ([]domain.Player, error)
	GetPlayerById(id uuid.UUID) (*domain.Player, error)
	CreateNewPlayer(player domain.Player) error
}

type PlayerServiceImpl struct {
	repo repository.PlayerRepository
}

func (s *PlayerServiceImpl) CreateNewPlayer(player domain.Player) error {
	if err := s.repo.CreateNewPlayer(player); err != nil {
		return err
	}
	return nil
}

func (s *PlayerServiceImpl) GetPlayerById(id uuid.UUID) (*domain.Player, error) {
	//TODO implement me
	panic("implement me")
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

//func (s *playerServiceImpl) CreateNewPlayer(player Player) error {
//	if err := s.repo.Create(player); err != nil {
//		return err
//	}
//	return nil
//}

func NewPlayerService(repo repository.PlayerRepository) *PlayerServiceImpl {
	return &PlayerServiceImpl{repo: repo}
}
