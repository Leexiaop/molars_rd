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
	PageNum    int
	PageSize   int
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

func (r *User) getMaps() map[string]interface{}  {
	maps := make(map[string]interface{})
	return maps
}