package controllers

import (
	"net/http"
	"secondhand/config"
	"secondhand/models"

	"github.com/gin-gonic/gin"
)

func GetSellers(c *gin.Context) {
	var sellers []models.Seller
	if err := config.DB.Preload("Status").Find(&sellers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sellers)
}
