package handler

import (
	"go_game/game/core"

	"gorm.io/gorm"
)

type gormPlayerRepository struct {
	db *gorm.DB
}

// Create implements core.PlayerRepository.
func (r *gormPlayerRepository) Create(player core.Player) error {
	if result := r.db.Create(&player); result.Error != nil {
		// Handle database errors
		return result.Error
	}
	return nil
}

// Delete implements core.PlayerRepository.
func (*gormPlayerRepository) Delete(player core.Player) error {
	panic("unimplemented")
}

// GetAll implements core.PlayerRepository.
func (r *gormPlayerRepository) GetAll(player core.Player) error {
	panic("unimplemented")
}

// GetById implements core.PlayerRepository.
func (*gormPlayerRepository) GetById(player core.Player) error {
	panic("unimplemented")
}

// Update implements core.PlayerRepository.
func (*gormPlayerRepository) Update(player core.Player) error {
	panic("unimplemented")
}

func NewGormPlayerRepository(db *gorm.DB) *gormPlayerRepository {
	return &gormPlayerRepository{db: db}
}
