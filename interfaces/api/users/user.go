package users

import (
	"net/http"

	"apiii/infrastructure/db"
	"apiii/usecases/users"
	"apiii/interfaces/repositories"
	"github.com/gin-gonic/gin"
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

func (controller *UserController)GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}