package handler

import (
	"github.com/google/uuid"
	"go_game/game/core"
	"gorm.io/gorm"
)

type GormPlayerRepository struct {
	db *gorm.DB
}

func (r *GormPlayerRepository) GetAllPlayer() ([]core.Player, error) {
	var player []core.Player
	if result := r.db.Find(&player); result.Error != nil {
		return nil, result.Error
	}

	return player, nil
}

func (r *GormPlayerRepository) GetPlayerById(id uuid.UUID) (*core.Player, error) {
	//TODO implement me
	panic("implement me")
}

func (r *GormPlayerRepository) CreateNewPlayer(player core.Player) error {
	if result := r.db.Create(&player); result.Error != nil {
		return result.Error
	}
	return nil

}

//// Create implements core.PlayerRepository.
//func (r *GormPlayerRepository) Create(player core.Player) error {
//	if result := r.db.Create(&player); result.Error != nil {
//		// Handle database errors
//		return result.Error
//	}
//	return nil
//}

func NewGormPlayerRepository(db *gorm.DB) core.PlayerRepository {
	return &GormPlayerRepository{db: db}
}
