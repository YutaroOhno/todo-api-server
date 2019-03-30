package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"apiii/interfaces/api/todos"
	"apiii/infrastructure/db"
)

func Run(db *db.DB) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		//とりあえずフロントからのアクセスを許可したいので、記述。
		AllowOrigins:     []string{"http://127.0.0.1:8080"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin"},
	}))

	// 全てのtodo取得
	todoController := todos.NewTodoController(db)
	router.GET("/todos", todoController.GetAllTodo)
	router.POST("/todos", todoController.CreateTodo)
	router.DELETE("/todos/:id", todoController.DeleteTodo)

	router.Run()
}