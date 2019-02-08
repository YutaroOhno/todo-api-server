package main

import (
	"apiii/infrastructure/server"
	"apiii/infrastructure/db"
)

func main() {
	var newDB db.Database
	newDB = db.NewMysql()
	db := newDB.Open()
	defer db.Close()

	server.Run(db)
}
