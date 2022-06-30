package models

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id" binding:"required"`
	Email     string `json:"email"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	Role      string `json:"role"`
	Position  string `json:"position"`
	Salary    uint   `json:"salary"`
	Teams     []Team `gorm:"many2many:user_teams;"`
}
