package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       int    `gotm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar string `json:"avatar"`
	Auth int `json:"auth"`
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
