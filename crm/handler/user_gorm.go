package handler

import (
	"go_game/crm/core/domain"
	"go_game/crm/core/repository"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repository.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) GetAllUser() ([]domain.User, error) {
	var user []domain.User
	if result := r.db.Find(&user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil

}

func (r *GormUserRepository) GetUserById(id string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *GormUserRepository) CreateNewUser(user domain.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}
