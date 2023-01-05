package common

import (
	// "gorm.io/driver/postgres"
	"gotodoapp/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteDB(dbUrl string) *gorm.DB {
	// db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln("INIT: error_db_initial_conn: ", err)
	}

	return db // db
}

func SetupDBTables(db *gorm.DB) {
	err := db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalln("INIT: error_db_initial_tables: ", err)
	}
}
