package models

type Team struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name" binding:"required"`
}
