package users

import (
	"apiii/infrastructure/db"
	"apiii/usecases/users"
	"apiii/interfaces/repositories"
)

type UserController struct {
	Usecase *users.UserUsecase
}

func NewUserController(db *db.DB) *UserController {
	return &UserController{
		Usecase: &users.UserUsecase{
			UserRepository: &repositories.UserRepository{},
		},
	}
}