package controllers

import "gorm.io/gorm"

type DBController struct {
	database *gorm.DB
}

func NewDBController(db *gorm.DB) *DBController {
	return &DBController{db}
}
