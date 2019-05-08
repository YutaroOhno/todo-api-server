package main

import (
	"apiii/infrastructure/db"
	"apiii/infrastructure/server"

	"apiii/infrastructure/logging"
	ulogging "apiii/usecases/logging"
)

func main() {
	var newDB db.Database
	newDB = db.NewMysql()
	db := newDB.Open()
	defer db.Close()


	// 環境によってlogginを分ける想定。現状ローカルのみ。
	var ulogging ulogging.Logging
	ulogging = logging.NewLogrusLogging()

	server.Run(db, ulogging)
}
