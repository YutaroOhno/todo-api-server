package todos

import (
	"net/http"
	"apiii/infrastructure/db"
	"apiii/usecases/todos"
	"apiii/usecases/ports"
	"apiii/interfaces/repositories"
	"github.com/gin-gonic/gin"
)

// json定義
type todoJSON struct {
	Title string `json:"title" binding:"required"`
	Text string `json:"text" binding:"required"`
}

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

func (controller *TodoController) CreateTodo(c *gin.Context) {
	var json todoJSON
	if err := c.BindJSON(&json); err != nil {
		//Errorの方はレスポンスオブジェクト作ってもいいかもね。
		c.JSON(http.StatusBadRequest, err.Error())
		return 
	}

	input := &ports.TodoInputPort{
		Title: json.Title,
		Text: json.Text,
	}
	
	output, err := controller.Usecase.CreateTodo(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}
