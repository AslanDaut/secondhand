package models

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Nickname     string `json:"nickname"`
	Mail         string `json:"mail"`
	PasswordHash string `json:"password_hash"`
}
