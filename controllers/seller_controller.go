package controllers

import (
	"net/http"
	"os"
	"secondhand/config"
	"secondhand/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// GetSellers функция для получения списка продавцов
func GetSellers(c *gin.Context) {
	var sellers []models.Seller
	if err := config.DB.Find(&sellers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sellers"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sellers": sellers})
}

type RegisterInput struct {
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	PhoneNumber   string `json:"phone_number" binding:"required"`
	StoreName     string `json:"store_name" binding:"required"`
	City          string `json:"city" binding:"required"`
	Address       string `json:"address"`
	InstagramLink string `json:"instagram_link"`
	TgChannelLink string `json:"tg_channel_link"`
	TgLink        string `json:"tg_link"`
	Password      string `json:"password" binding:"required"`
}

type LoginInput struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

// Регистрация продавца
func RegisterSeller(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Создание нового продавца
	seller := models.Seller{
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		PhoneNumber:   input.PhoneNumber,
		StoreName:     input.StoreName,
		City:          input.City,
		Address:       input.Address,
		InstagramLink: input.InstagramLink,
		TgChannelLink: input.TgChannelLink,
		TgLink:        input.TgLink,
		StatusID:      0,                      // Устанавливаем статус "на модерации"
		Password:      string(hashedPassword), // Сохраняем хешированный пароль
	}

	if err := config.DB.Create(&seller).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

// Логин продавца
func LoginSeller(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Поиск продавца по номеру телефона
	var seller models.Seller
	if err := config.DB.Where("phone_number = ?", input.PhoneNumber).First(&seller).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid phone number or password"})
		return
	}

	// Проверка пароля
	err := bcrypt.CompareHashAndPassword([]byte(seller.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid phone number or password"})
		return
	}

	// Создание JWT токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": seller.ID,
		"status":  0, // Статус "на модерации"
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
