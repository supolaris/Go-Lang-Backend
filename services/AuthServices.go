package services

import (
	"errors"
	internal "goProject/internals/models"

	"gorm.io/gorm"
)

type AuthService struct{
	db *gorm.DB
}

func InitAuthService(database *gorm.DB) *AuthService {
database.AutoMigrate(&internal.User{})
	return &AuthService{
		db: database,
	}
}

func (n *AuthService) Register(email *string, password *string)(*internal.User, error){
	if email == nil {
		return nil, errors.New("Email is required")
	}
	if password == nil {
		return nil, errors.New("Password is required")
	}
	var user internal.User
	user.Email = *email
	user.Password = *password
	if err := n.db.Create(&user).Error; err != nil {
		return nil,err
	}
	return &user,nil;
}

func (n *AuthService) Login(email *string, password *string)(*internal.User, error){
	if email == nil {
		return nil, errors.New("Email is required")
	}
	if password == nil {
		return nil, errors.New("Password is required")
	}
	var user internal.User
	if err := n.db.Where("email = ?", email).Where("password = ?",password).Find(&user).Error; err !=nil{
		return nil,err
	}
	if user.Email == ""{
		return nil,errors.New("No user found")
	}
	return &user, nil;
}