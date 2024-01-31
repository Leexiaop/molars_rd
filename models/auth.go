package models

import (
	"time"

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
	err := db.Where(User{Username: username, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

func AddUser (username, password, auth string) error {
	user := User{
		Username: username,
		Password: password,
		Auth: auth,
		CreatedBy: username,
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifieldOn", time.Now().Unix())
	return nil
}