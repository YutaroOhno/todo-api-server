package repositories

import (
	"apiii/entities"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	SelectByUserID(*gorm.DB, uint) (*entities.User, error)
}