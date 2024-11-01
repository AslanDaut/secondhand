package models

type Review struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserID   uint   `json:"user_id"`
	User     User   `json:"user" gorm:"foreignKey:UserID"`
	SellerID uint   `json:"seller_id"`
	Seller   Seller `json:"seller" gorm:"foreignKey:SellerID"`
}
