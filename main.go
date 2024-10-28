package main

import (
	"secondhand/config"
	"secondhand/models" // Импортируем пакет models
	"secondhand/routes"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Подключение к базе данных
	config.ConnectDatabase()

	// Пример использования структур из models
	// Создание экземпляров структур и вывод для проверки импорта
	product := models.Product{
		ProductName: "Example Product",
		Price:       99.99,
	}
	fmt.Printf("Product: %s, Price: %.2f\n", product.ProductName, product.Price)

	user := models.User{
		Nickname: "example_user",
		Mail:     "user@example.com",
	}
	fmt.Printf("User: %s, Email: %s\n", user.Nickname, user.Mail)

	// Настройка маршрутов
	routes.SetupRoutes(r)

	// Запуск сервера
	r.Run() // Слушаем и обслуживаем на 0.0.0.0:8080
}
