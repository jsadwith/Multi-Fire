/*

database.go

Database handling
Implements gorm ORM: https://github.com/jinzhu/gorm

*/

package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func OpenDbConn() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./database/data.db")
	if err != nil {
		panic("Failed to connect to db")
	}

	return db
}
