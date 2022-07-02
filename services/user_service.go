package services

import (
	"sgcu65/models"
	"sgcu65/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	AddUser(user models.User) (models.User, error)
	GetUser(userId string) (models.User, error)
	DeleteUser(userId string) error
}

func NewUserService(r repository.UserRepository) UserService {
	return userService{
		userRepository: r,
	}
}

func (u userService) AddUser(user models.User) (models.User, error) {
	return u.userRepository.AddUser(user)
}

func (u userService) GetUser(userId string) (models.User, error) {
	return u.userRepository.GetUser(userId)
}

func (u userService) DeleteUser(userId string) error {
	return u.userRepository.DeleteUser(userId)
}
