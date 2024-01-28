package basic

import (
	"net/http"

	"github.com/Leexiaop/molars_rd/pkg/app"
	"github.com/Leexiaop/molars_rd/pkg/e"
	"github.com/Leexiaop/molars_rd/pkg/setting"
	"github.com/Leexiaop/molars_rd/pkg/util"
	"github.com/Leexiaop/molars_rd/service/user_service"
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