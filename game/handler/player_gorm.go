package handler

import (
	"go_game/game/core/domain"
	"go_game/game/core/repository"
	"gorm.io/gorm"
)

type GormPlayerRepository struct {
	db *gorm.DB
}

func NewGormPlayerRepository(db *gorm.DB) repository.PlayerRepository {
	return &GormPlayerRepository{db: db}
}

func (r *GormPlayerRepository) GetAllPlayer() ([]domain.Player, error) {
	var player []domain.Player
	if result := r.db.Find(&player); result.Error != nil {
		return nil, result.Error
	}
	return player, nil
}

func (r *GormPlayerRepository) GetPlayerById(id string) (*domain.Player, error) {
	var player domain.Player
	if result := r.db.First(&player, "player_id", id); result.Error != nil {
		return nil, result.Error
	}
	return &player, nil
}

func (r *GormPlayerRepository) CreateNewPlayer(player domain.Player) error {
	if result := r.db.Create(&player); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormPlayerRepository) UpdatePlayer(player domain.Player) (*domain.Player, error) {
	var playerDB domain.Player
	if result := r.db.First(&playerDB, "player_id", player.PlayerID); result.Error != nil {
		return nil, result.Error
	}
	playerDB.PlayerName = player.PlayerName

	return &playerDB, nil
}

func (r *GormPlayerRepository) DeletePlayer(id string) error {
	var player domain.Player
	if result := r.db.Delete(&player, "player_id", id); result.Error != nil {
		return result.Error
	}
	return nil
}
