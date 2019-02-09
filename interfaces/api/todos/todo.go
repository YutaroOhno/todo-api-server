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
			DB: db,
		},
	}
}

func (controller *TodoController) GetAllTodo(c *gin.Context) {

	todos, err := controller.Usecase.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": todos,
	})
}
