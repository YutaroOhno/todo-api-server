package db

import (
	"github.com/jinzhu/gorm"
)

type DB struct {
	gormDB *gorm.DB
}

type Database interface {
	Open() *DB
}

func (db *DB) Close() {
	db.gormDB.Close()
}