package routes

import (
	"secondhand/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", controllers.RegisterSeller)
		api.POST("/login", controllers.LoginSeller)
		api.GET("/sellers", controllers.GetSellers)
	}
}
