package models

type Seller struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	PhoneNumber   string `json:"phone_number"`
	StoreName     string `json:"store_name"`
	City          string `json:"city"`
	Address       string `json:"address"`
	InstagramLink string `json:"instagram_link"`
	TgChannelLink string `json:"tg_channel_link"`
	TgLink        string `json:"tg_link"`
	StatusID      uint   `json:"status_id"`
	Status        Status `json:"status" gorm:"foreignKey:StatusID"`
}
