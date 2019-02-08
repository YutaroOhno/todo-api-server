package repositories

import (
	"apiii/entities"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
}

func (repo *UserRepository) SelectByUserID(db *gorm.DB, userId uint) (*entities.User, error) {
	return &entities.User{}, nil
}