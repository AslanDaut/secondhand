package config

import (
	"fmt"
	"log"
	"os"
	"secondhand/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// Загружаем .env файл
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env")
	}
}

func ConnectDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных: ", err)
	}

	// Автоматическая миграция схемы базы данных
	database.AutoMigrate(&models.Seller{}, &models.Status{}, &models.Product{}, &models.Category{}, &models.User{}, &models.Review{})

	DB = database
}
