package services

import (
	"github.com/duvrdx/oasys-server/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) CreateUser(user *models.UserInput) (*models.User, error) {
	u := user.ToUser()
	return u, s.DB.Create(&u).Error
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) UpdateUser(id string, userInput *models.UserInput) (*models.User, error) {
	var userToUpdate models.User
	if err := s.DB.Where("id = ?", id).First(&userToUpdate).Error; err != nil {
		return nil, err
	}

	updatedUser, err := userInput.ToUser().Update(&userToUpdate)
	if err != nil {
		return nil, err
	}

	return updatedUser, s.DB.Save(&updatedUser).Error
}

func (s *UserService) DeleteUser(id string) error {
	return s.DB.Delete(&models.User{}, id).Error
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
