package model

import (
	"github.com/msojocs/AutoTask/v1/db"
)

type User struct {
	ID       int64
	Login    string `form:"login"`
	Name     string `form:"name"`
	Password string `form:"password"`
	Email    string
	Nick     string
	Status   int
	Avatar   string
}

func (User) TableName() string {
	return "at_users"
}

func (user *User) Save() int64 {
	db.DB.Create(&user)
	return user.ID
}

func (user *User) CheckPassword(pass string) (bool, error) {
	return true, nil
}

// GetUserByEmail 用Email获取用户
func GetUserByEmail(email string) (User, error) {
	var user User
	result := db.DB.Set("gorm:auto_preload", true).Where("email = ?", email).First(&user)
	return user, result.Error
}
