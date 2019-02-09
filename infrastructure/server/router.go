package server

import (
	"github.com/gin-gonic/gin"
	"apiii/interfaces/api/users"
	"apiii/infrastructure/db"
)

func Run(db *db.DB) {
	router := gin.Default()

	userController := users.NewUserController(db)
	router.GET("/users", userController.GetUser)

	router.Run()
}