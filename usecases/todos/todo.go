package todos

import (
	"apiii/entities"
	"apiii/infrastructure/db"
	"apiii/usecases/ports"
	"apiii/usecases/repositories"
	"apiii/usecases"
	"apiii/usecases/logging"
)

type TodoUsecase struct {
	TodoRepository repositories.TodoRepository
	DB             *db.DB
	Logging 		logging.Logging
}

func (usecase *TodoUsecase) GetAllTodo() ([]entities.Todo, *usecases.UError) {
	todos, err := usecase.TodoRepository.FindAll(usecase.DB.GormDB)
	if uerr := usecases.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return nil, uerr
	}

	return todos, nil
}

func (usecase *TodoUsecase) CreateTodo(input *ports.TodoInputPort) (*ports.TodoOutputPort, error) {
	todo := &entities.Todo{
		Title: input.Title,
		Text:  input.Text,
	}

	if err := usecase.TodoRepository.Insert(usecase.DB.GormDB, todo); err != nil {
		return nil, err
	}

	return createOutputPort(todo), nil
}

func (usecase *TodoUsecase) DeleteTodo(id int) error {
	todo := entities.Todo{}
	todo.ID = id
	if err := usecase.TodoRepository.Delete(usecase.DB.GormDB, &todo); err != nil {
		return err
	}

	return nil
}

func createOutputPort(todo *entities.Todo) *ports.TodoOutputPort {
	return &ports.TodoOutputPort{
		Title: todo.Title,
		Text:  todo.Text,
	}
}
