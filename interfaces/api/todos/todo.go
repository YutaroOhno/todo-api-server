package todos

import (
	"apiii/infrastructure/db"
	"apiii/interfaces/repositories"
	"apiii/usecases/ports"
	"apiii/usecases/todos"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// json定義
type todoJSON struct {
	Title string `form:"title" binding:"required"`
	Text  string `form:"text" json:"text" binding:"required"`
}

type TodoController struct {
	Usecase *todos.TodoUsecase
}

func NewTodoController(db *db.DB) *TodoController {
	return &TodoController{
		Usecase: &todos.TodoUsecase{
			TodoRepository: &repositories.TodoRepository{},
			DB:             db,
		},
	}
}

func (controller *TodoController) GetAllTodo(c *gin.Context) {
	todos, err := controller.Usecase.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func (controller *TodoController) CreateTodo(c *gin.Context) {
	var json todoJSON

	if err := c.Bind(&json); err != nil {
		//Errorの方はレスポンスオブジェクト作ってもいいかもね。
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	input := &ports.TodoInputPort{
		Title: json.Title,
		Text:  json.Text,
	}

	output, err := controller.Usecase.CreateTodo(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}

func (controller *TodoController) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = controller.Usecase.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
}
