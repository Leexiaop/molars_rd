package basic

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Leexiaop/molars_rd/models"
	"github.com/Leexiaop/molars_rd/pkg/app"
	"github.com/Leexiaop/molars_rd/pkg/e"
	"github.com/Leexiaop/molars_rd/pkg/setting"
	"github.com/Leexiaop/molars_rd/pkg/util"
	"github.com/Leexiaop/molars_rd/service/user_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	userServeice := user_service.User{
		PageNum: util.GetPage(ctx),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := userServeice.Counts()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_USER_LIST_FAIL, nil)
		return
	}

	users, err := userServeice.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_USER_LIST_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list": users,
		"total": total,
	})
}

func EditUsers(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	jsonData, _ := ctx.GetRawData()

	var m models.User
	json.Unmarshal(jsonData, &m)
	id := m.ID
	username := m.Username
	password := m.Password
	avatar := m.Avatar
	phone := m.Phone
	auth := m.Auth

	fmt.Print("auth===", auth)

	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不能为空!")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	var userService = user_service.User{}
	if id > 0 {
		userService.ID = id
	}
	if username != "" {
		userService.Username = username
	}
	if password != "" {
		userService.Password = password
	}
	if auth != "" {
		userService.Auth = auth
	}

	if avatar != "" {
		userService.Avatar = avatar
	}
	if phone != "" {
		userService.Phone = phone
	}

	exists, err := userService.ExistById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_USER_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	user, err := userService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_RECORD_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user)
}
