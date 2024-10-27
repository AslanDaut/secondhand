package models

type Product struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	SellerID    uint     `json:"seller_id"`
	Seller      Seller   `json:"seller" gorm:"foreignKey:SellerID"`
	Photo1      string   `json:"photo1"`
	Photo2      string   `json:"photo2"`
	Photo3      string   `json:"photo3"`
	Photo4      string   `json:"photo4"`
	Photo5      string   `json:"photo5"`
	ProductName string   `json:"product_name"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
	Brand       string   `json:"brand"`
	State       string   `json:"state"`
	Price       float64  `json:"price"`
	Size        string   `json:"size"`
	Defect      string   `json:"defect"`
	Description string   `json:"description"`
}
