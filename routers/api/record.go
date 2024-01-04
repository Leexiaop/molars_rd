package api

import (
	"fmt"
	"net/http"

	"github.com/Leexiaop/molars_rd/models"
	"github.com/Leexiaop/molars_rd/pkg/e"
	"github.com/Leexiaop/molars_rd/pkg/setting"
	"github.com/Leexiaop/molars_rd/pkg/util"
	"github.com/gin-gonic/gin"
)

//	获取产品列表

func GetRecords (ctx * gin.Context) {
	productId := ctx.Query("product_id")

	fmt.Printf(productId, 898989)

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if len(productId) != 0 && productId != "" {
		maps["product_id"] = productId
	}

	code := e.SUCCESS

	data["list"] = models.GetRecords(util.GetPage(ctx), setting.PageSize, maps)
	data["total"] = models.GetRecordsTotal(maps)

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}