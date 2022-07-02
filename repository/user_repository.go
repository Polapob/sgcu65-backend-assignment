package repository

import (
	"sgcu65/models"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	AddUser(user models.User) (models.User, error)
	GetUser(userId string) (models.User, error)
	DeleteUser(userId string) error
	Migrate() error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) Migrate() error {
	return u.DB.AutoMigrate(&models.User{})
}

func (u userRepository) AddUser(user models.User) (models.User, error) {
	err := u.DB.Create(&user).Error
	return user, err
}

func (u userRepository) GetUser(userId string) (models.User, error) {
	var user models.User
	err := u.DB.Where("id = ?", userId).First(&user).Error
	return user, err
}

func (u userRepository) DeleteUser(userId string) error {
	err := u.DB.Delete(&models.User{}, userId).Error
	return err
}
