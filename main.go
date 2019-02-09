package main

import (
	"apiii/infrastructure/server"
	"apiii/infrastructure/db"
	"log"
)

func main() {
	var newDB db.Database
	newDB = db.NewMysql()
	db := newDB.Open()
	log.Printf("aaaaaaaaaaaa")
	log.Printf("%v", db)
	defer db.Close()

	server.Run(db)
}
