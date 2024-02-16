package handler

import (
	"github.com/google/uuid"
	"go_game/game/core/domain"
	"go_game/game/core/repository"
	"gorm.io/gorm"
)

type GormPlayerRepository struct {
	db *gorm.DB
}

func (r *GormPlayerRepository) GetAllPlayer() ([]domain.Player, error) {
	var player []domain.Player
	if result := r.db.Find(&player); result.Error != nil {
		return nil, result.Error
	}

	return player, nil
}

func (r *GormPlayerRepository) GetPlayerById(id uuid.UUID) (*domain.Player, error) {
	//TODO implement me
	panic("implement me")
}

func (r *GormPlayerRepository) CreateNewPlayer(player domain.Player) error {
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

func NewGormPlayerRepository(db *gorm.DB) repository.PlayerRepository {
	return &GormPlayerRepository{db: db}
}
