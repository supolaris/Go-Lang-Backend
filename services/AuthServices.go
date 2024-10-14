package services

import (
	"errors"
	internal "goProject/internals/models"
	"goProject/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func InitAuthService(database *gorm.DB) *AuthService {
	database.AutoMigrate(&internal.User{})
	return &AuthService{
		db: database,
	}
}

func (n *AuthService) CheckUserExists(email *string) bool {
	var user internal.User
	if err := n.db.Where("email = ?", email).Find(&user); err != nil {
		return false
	}
	if user.Email != "" {
		return true
	}
	return false
}

func (n *AuthService) Register(email *string, password *string) (*internal.User, error) {
	if email == nil {
		return nil, errors.New("email is required")
	}
	if password == nil {
		return nil, errors.New("password is required")
	}
	hashedPwd, err := utils.HasPassword(*password)
	if err != nil {
		return nil, err
	}
	var user internal.User
	user.Email = *email
	user.Password = hashedPwd
	if err := n.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (n *AuthService) Login(email *string, password *string) (*internal.User, error) {
	if email == nil {
		return nil, errors.New("email is required")
	}
	if password == nil {
		return nil, errors.New("password is required")
	}
	var user internal.User
	if err := n.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}

	if user.Email == "" {
		return nil, errors.New("user not found")
	}

	if !utils.CheckPasswordHash(*password, user.Password) {
		return nil, errors.New("incorrect password")
	}
	return &user, nil
}
