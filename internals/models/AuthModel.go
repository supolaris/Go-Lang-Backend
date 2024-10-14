package model

type User struct {
	Id       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Email    string `gorm:"unique;not null" json:"email" binding:"required"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "User"
}
