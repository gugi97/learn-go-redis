package db

import (
	"fmt"
	"github.com/gugi97/learn-go-redis/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitPostgres() *gorm.DB {
	username := "admin"
	password := "admin"
	host := "localhost"
	port := "5432"
	database := "learn"

	dbURL := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s`, username, password, host, port, database)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})
	return db
}
