package core

import (
	"github.com/google/uuid"
)

type PlayerService interface {
	GetAllPlayer() ([]Player, error)
	GetPlayerById(id uuid.UUID) (*Player, error)
	CreateNewPlayer(player Player) error
}

type PlayerServiceImpl struct {
	repo PlayerRepository
}

func (s *PlayerServiceImpl) CreateNewPlayer(player Player) error {
	if err := s.repo.CreateNewPlayer(player); err != nil {
		return err
	}
	return nil
}

func (s *PlayerServiceImpl) GetPlayerById(id uuid.UUID) (*Player, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PlayerServiceImpl) GetAllPlayer() ([]Player, error) {
	player, err := s.repo.GetAllPlayer()
	if err != nil {
		return nil, err
	}

	var playerResponses []Player

	for _, player := range player {
		response := Player{
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

func NewPlayerService(repo PlayerRepository) *PlayerServiceImpl {
	return &PlayerServiceImpl{repo: repo}
}
