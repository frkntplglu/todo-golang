package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	return gorm.Open(postgres.Open("host=localhost user=postgres password=fuki123 dbname=planner port=5432 sslmode=disable"), &gorm.Config{})
}
