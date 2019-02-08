package users

import (
	"apiii/usecases/repositories"
)

type UserUsecase struct {
	UserRepository repositories.UserRepository
}