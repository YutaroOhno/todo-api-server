package todos

import (
	"net/http"

	"apiii/infrastructure/db"
	"apiii/usecases/todos"
	"apiii/interfaces/repositories"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	Usecase *todos.TodoUsecase
}

func NewTodoController(db *db.DB) *TodoController {
	return &TodoController{
		Usecase: &todos.TodoUsecase{
			TodoRepository: &repositories.TodoRepository{},
		},
	}
}

func (controller *TodoController) GetAllTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// func (controller *UserController)GetUser(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "pong",
// 	})
// }