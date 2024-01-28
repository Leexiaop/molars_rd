package util

import "github.com/Leexiaop/molars_rd/pkg/setting"

func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}