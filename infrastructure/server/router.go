package server

import (
	"github.com/gin-gonic/gin"
	"apiii/interfaces/api/todos"
	"apiii/infrastructure/db"
)

func Run(db *db.DB) {
	router := gin.Default()

	todoController := todos.NewTodoController(db)
	router.GET("/todos", todoController.GetAllTodo)
	router.Run()
}