package server

import (
	"github.com/gin-gonic/gin"
	"apiii/interfaces/api/users"
	"apiii/infrastructure/db"
	"fmt"
)

func Run(db *db.DB) {
	router := gin.Default()

	userController := users.NewUserController(db)
	fmt.Println(userController)

	router.Run()
}