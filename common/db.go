package common

import (
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	"gotodoapp/models"
)

func NewPostgresDB(dbUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
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
