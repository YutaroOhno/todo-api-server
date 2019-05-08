package server

import (
	"apiii/infrastructure/db"
	"apiii/interfaces/api/todos"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"apiii/usecases/logging"
)

func Run(db *db.DB, logging logging.Logging) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		//とりあえずフロントからのアクセスを許可したいので、記述。
		AllowOrigins: []string{"http://127.0.0.1:8080"},
		AllowMethods: []string{"GET", "PUT", "PATCH", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin"},
	}))

	// 全てのtodo取得
	todoController := todos.NewTodoController(db, logging)
	router.GET("/todos", todoController.GetAllTodo)
	router.POST("/todos", todoController.CreateTodo)
	router.DELETE("/todos/:id", todoController.DeleteTodo)

	router.Run()
}
