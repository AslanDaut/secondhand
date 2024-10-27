package routes

import (
	"your-project/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/sellers", controllers.GetSellers)
		api.GET("/products", controllers.GetProducts)
		// Добавьте дополнительные маршруты по необходимости
	}
}
