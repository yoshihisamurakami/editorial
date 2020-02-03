package db

import (
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("postgres", "user=" + os.Getenv("DBUSERNAME") + " dbname=news sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetDb() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
