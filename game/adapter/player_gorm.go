package adapter

import (
	"go_game/game/core"

	"gorm.io/gorm"
)

type gormPlayerRepository struct {
	db *gorm.DB
}

// Create implements core.PlayerRepository.
func (*gormPlayerRepository) Create(player core.Player) error {
	panic("unimplemented")
}

// Delete implements core.PlayerRepository.
func (*gormPlayerRepository) Delete(player core.Player) error {
	panic("unimplemented")
}

// Get implements core.PlayerRepository.
func (*gormPlayerRepository) Get(player core.Player) error {
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

func NewGormPlayerRepository(db *gorm.DB) core.PlayerRepository {
	return &gormPlayerRepository{db: db}
}

func (r *gormPlayerRepository) Save(player core.Player) error {
	if result := r.db.Create(&player); result.Error != nil {
		// Handle database errors
		return result.Error
	}
	return nil
}
