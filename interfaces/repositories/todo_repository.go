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

func (repo *TodoRepository) Insert(db *gorm.DB, todo *entities.Todo) error {
	if err := db.Create(todo).Error; err != nil {
		return err
	}

	return nil
}

func (repo *TodoRepository) Delete(db *gorm.DB, todo *entities.Todo) error {
	if err := db.Delete(todo).Error; err != nil {
		return err
	}

	return nil
}
