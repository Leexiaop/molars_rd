package user_service

import (
	"encoding/json"

	"github.com/Leexiaop/molars_rd/models"
	"github.com/Leexiaop/molars_rd/pkg/gredis"
	"github.com/Leexiaop/molars_rd/pkg/logging"
	"github.com/Leexiaop/molars_rd/service/cache_service"
)

type User struct {
	ID       int
	Username string
	Avatar   string
	Auth     string
	Phone    string
	Password string
	PageNum    int
	PageSize   int
	ModifieldBy string
}

func (u *User) Counts() (int, error) {
	return models.GetUserTotal(u.getMaps())
}
func (u *User) GetAll() ([]models.User, error) {
	var (
		users, cacheUsers []models.User
	)

	cache := cache_service.User{
		PageSize:  u.PageSize,
		PageNum:   u.PageNum,
	}
	key := cache.GetUsersKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheUsers)
			return cacheUsers, nil
		}
	}
	users, err := models.GetAllUser(u.PageNum, u.PageSize, u.getMaps())
	if err != nil {
		return nil, err
	}
	gredis.Set(key, users, 3600)
	return users, nil
}

func (u * User) ExistById () (bool, error) {
	return models.ExistUserId(u.ID)
}

func (u *User) getMaps() map[string]interface{}  {
	maps := make(map[string]interface{})
	if u.Username != "" {
		maps["username"] = u.Username
	}
	if u.Password != "" {
		maps["password"] = u.Password
	}
	if u.Avatar != "" {
		maps["avatar"] = u.Avatar
	}
	if u.Phone != "" {
		maps["phone"] = u.Phone
	}
	if u.Auth != "" {
		maps["auth"] = u.Auth
	}
	return maps
}

func (u *User) Edit() (*models.User, error) {
	data := make(map[string]interface{})
	if u.Username != "" {
		data["username"] = u.Username
	}
	if u.Password != "" {
		data["password"] = u.Password
	}
	if u.Avatar != "" {
		data["avatar"] = u.Avatar
	}
	if u.Phone != "" {
		data["phone"] = u.Phone
	}
	if u.Auth != "" {
		data["auth"] = u.Auth
	}
	if (u.ModifieldBy != "") {
		data["modifield_by"] = u.ModifieldBy
	}

	return models.EditUser(u.ID, data)
}