package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var sqliteDB *gorm.DB

// InitSqliteDB initialize sqlite database
func InitSqliteDB() {
	dbName := os.Getenv("DB_NAME")
	if os.Getenv("ENVIRONMENT") != "production" {
		dbName = dbName + "-" + os.Getenv("ENVIRONMENT")
	}
	dbName = dbName + ".db"

	fmt.Println("Connecting to sqlite database: ", dbName)
	db, err := gorm.Open(sqlite.Open(dbName))

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	sqliteDB = db
}

func GetSqliteDB() *gorm.DB {
	if sqliteDB == nil {
		fmt.Println("sqliteDB is nil")
		InitSqliteDB()
	}
	return sqliteDB
}
