package repositories

import (
	"apiii/entities"
	"github.com/jinzhu/gorm"
)

type TodoRepository interface {
	SelectByID(*gorm.DB, uint) (*entities.Todo, error)
}