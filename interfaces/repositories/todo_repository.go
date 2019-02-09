package repositories

import (
	"apiii/entities"
	"github.com/jinzhu/gorm"
)

type TodoRepository struct {
}

func (repo *TodoRepository) SelectByID(db *gorm.DB, userId uint) (*entities.Todo, error) {
	return &entities.Todo{}, nil
}