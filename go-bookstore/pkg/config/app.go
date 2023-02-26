package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("mysql", "ujjwal:password/simplerest?charset=utf8&parseTime=True&loc=Local")
	//d, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	//d, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	d, err := gorm.Open("sqlite3", "database.db")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
