package repositories

import (
	"apiii/entities"
	"github.com/jinzhu/gorm"
)

type TodoRepository interface {
	FindAll(*gorm.DB) ([]entities.Todo, error)
}