package models

import "time"

type Task struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	Status   bool      `json:"status"`
	Deadline time.Time `json:"deadline"`
	Teams    []Team    `gorm:"many2many:task_teams;"`
	Users    []User    `gorm:"many2many:task_users;"`
}
