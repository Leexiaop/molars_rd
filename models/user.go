package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       int    `gotm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar string `json:"avatar"`
	Auth string `json:"auth"`
	Phone string `json:"phone"`
}
func GetUserTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&User{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetAllUser(pageNum int, pageSize int, maps interface{}) ([]User, error) {
	var (
		users []User
		err   error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Select([]string{"id", "username", "phone", "avatar", "auth"}).Where(maps).Find(&users).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Select([]string{"id", "username", "phone", "avatar", "auth"}).Where(maps).Find(&users).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

func ExistUserId (id int) (bool, error) {
	var user User
	err := db.Select("id").Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func EditUser (id int, data interface{}) (*User, error) {
	var (
		user User
		err error
	)

	err = db.Model(&user).Where("id = ?", id).Updates(data).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	err = db.Select([]string{"id", "username", "phone", "avatar", "auth"}).Find(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	
	return &user, nil
}
