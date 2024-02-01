package api

import (
	"encoding/json"
	"net/http"

	"github.com/Leexiaop/molars_rd/models"
	"github.com/Leexiaop/molars_rd/pkg/app"
	"github.com/Leexiaop/molars_rd/pkg/e"
	"github.com/Leexiaop/molars_rd/pkg/util"
	"github.com/Leexiaop/molars_rd/service/auth_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonData, _ := c.GetRawData()

	var m models.User

	json.Unmarshal(jsonData, &m)
	username := m.Username
	password := m.Password

	valid := validation.Validation{}
	valid.Required(username, "username").Message("用户名不能为空")
	valid.Required(password, "password").Message("密码不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.User{Username: username, Password: password}

	isHas, err := authService.CheckName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if !isHas {
		AddUser(c, username, password)
		return
	} else {
		isExist, err := authService.Check()
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			return
		}
		if !isExist {
			appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
			return
		}
		CreateToken(c, username, password)
	}
}

func AddUser(ctx *gin.Context, username, password string) {
	appG := app.Gin{C: ctx}
	auth := "0"

	valid := validation.Validation{}
	valid.Required(username, "username").Message("用户名不能为空!")
	valid.Required(password, "password").Message("密码不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.User{Username: username, Password: password, Auth: auth}
	

	err := authService.Add()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}

	// appG.Response(http.StatusOK, e.SUCCESS, nil)
	CreateToken(ctx, username, password)
	
}

func CreateToken(c *gin.Context, username, password string) {
	appG := app.Gin{C: c}
	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	user, _ := models.GetAuth(username, password)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]any{
		"token": token,
		"username": user.Username,
		"auth": user.Auth,
		"phone": user.Phone,
		"id": user.ID,
		"avatar": user.Avatar,
	})
}