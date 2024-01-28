package auth_service

import "github.com/Leexiaop/molars_rd/models"

type User struct {
	Username    string
	Password string
}

func (a *User) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}

