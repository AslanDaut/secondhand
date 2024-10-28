package routes

import (
	"secondhand/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/sellers", controllers.GetSellers)
		api.GET("/products", controllers.GetSellers)
		// Добавьте дополнительные маршруты по необходимости
	}
}
