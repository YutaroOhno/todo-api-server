package repositories

import (
	"apiii/entities"
	"github.com/jinzhu/gorm"
)

type TodoRepository interface {
	FindAll(*gorm.DB) ([]entities.Todo, error)
	Insert(*gorm.DB, *entities.Todo) error
	Delete(*gorm.DB, *entities.Todo) error
}