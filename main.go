package main

import (
	"your-project/config"
	"your-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Подключение к базе данных
	config.ConnectDatabase()

	// Настройка маршрутов
	routes.SetupRoutes(r)

	// Запуск сервера
	r.Run() // Слушаем и обслуживаем на 0.0.0.0:8080
}
