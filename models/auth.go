package models

import (
	"github.com/jinzhu/gorm"
)

func CheckAuth(username, password string) (bool, error) {
	var user User
	err := db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func GetAuth (username, password string)(*User, error) {
	var user User
	err := db.Select([]string{"id", "username", "avatar", "auth", "phone"}).Where(User{Username: username, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}