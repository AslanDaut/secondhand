package models

type Status struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	StatusName string `json:"status_name"`
	StatusDesc string `json:"status_desc"`
}
