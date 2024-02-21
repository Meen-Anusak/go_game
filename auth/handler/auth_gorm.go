package handler

import (
	"go_game/auth/core/domain"
	"go_game/auth/core/repository"
	"gorm.io/gorm"
)

type GormAuthRepository struct {
	db *gorm.DB
}

func NewGormAuthRepository(db *gorm.DB) repository.AuthRepository {
	return &GormAuthRepository{db: db}
}

func (r *GormAuthRepository) Login(login domain.Login) (*domain.User, error) {
	var user domain.User
	if result := r.db.Find(&user, "user_name", login.Username); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *GormAuthRepository) Register(user domain.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil

}
