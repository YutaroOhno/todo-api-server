package todos

import (
	"apiii/usecases/repositories"
)

type TodoUsecase struct {
	TodoRepository repositories.TodoRepository
}