package models

type Category struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
}
