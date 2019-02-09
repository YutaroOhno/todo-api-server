package todos

import (
	"apiii/entities"
	"apiii/usecases/repositories"
	"apiii/infrastructure/db"
)

type TodoUsecase struct {
	TodoRepository repositories.TodoRepository
	DB *db.DB
}

func(usecase *TodoUsecase) GetAllTodo() ([]entities.Todo, error){
	todos, err := usecase.TodoRepository.FindAll(usecase.DB.GormDB)
	if err != nil {
		panic(err.Error())
	}
	return todos, nil
}
