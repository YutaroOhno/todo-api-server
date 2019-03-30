package main

import (
	"apiii/infrastructure/db"
	"apiii/infrastructure/server"
)

func main() {
	var newDB db.Database
	newDB = db.NewMysql()
	db := newDB.Open()
	defer db.Close()

	server.Run(db)
}
