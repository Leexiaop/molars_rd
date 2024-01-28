package cache_service

import (
	"strconv"
	"strings"

	"github.com/Leexiaop/molars_rd/pkg/e"
)

type User struct {
	ID       int
	PageSize int
	PageNum  int
}

func (u *User) GetUserKey() string {
	return e.CACHE_USER + "_" + strconv.Itoa(u.ID)
}

func (u *User) GetUsersKey() string {
	keys := []string{
		e.CACHE_USER,
		"LIST",
	}
	if u.PageNum > 0 {
		keys = append(keys, strconv.Itoa(u.PageNum))
	}
	if u.PageSize > 0 {
		keys = append(keys, strconv.Itoa(u.PageSize))
	}
	return strings.Join(keys, "_")
}