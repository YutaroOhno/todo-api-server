package repositories

import (
	"apiii/entities"
	"github.com/jinzhu/gorm"
)

type TodoRepository struct {
}

func (repo *TodoRepository) FindAll(db *gorm.DB) ([]entities.Todo, error) {
	var todos []entities.Todo
	db.Find(&todos)
	return todos, nil
}