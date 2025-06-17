package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgresGORM() {
	dsn := "host=localhost user=postgres password=2040 dbname=to-do-app port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database: " + err.Error())
	}
	fmt.Println("Connected to PostgreSQL with GORM")
	DB = db
}
