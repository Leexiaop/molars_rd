package api

import (
	"encoding/json"
	"net/http"

	"github.com/Leexiaop/molars_rd/models"
	"github.com/Leexiaop/molars_rd/pkg/e"
	"github.com/Leexiaop/molars_rd/pkg/setting"
	"github.com/Leexiaop/molars_rd/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/astaxie/beego/validation"
)

//	获取记录列表
func GetRecords (ctx * gin.Context) {
	productId := ctx.Query("product_id")

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

func AddRecords (ctx * gin.Context) {
	jsonData, _ := ctx.GetRawData()

	var m models.Record
	json.Unmarshal(jsonData, &m)

	price := m.Price
	count := m.Count
	url := m.Url
	productId := m.ProductId
	
	valid := validation.Validation{}
	valid.Required(price, "price").Message("单价不能为空!")
	valid.Required(count, "count").Message("数量不能为空")
	valid.Required(productId, "productId").Message("产品ID不能为空")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		models.AddRecords(price, count, productId, url)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}
func EditRecords (ctx * gin.Context) {
	jsonData, _ := ctx.GetRawData()

	var m models.Record
	json.Unmarshal(jsonData, &m)
	id := m.ID
	price := m.Price
	count := m.Count
	url := m.Url
	productId := m.ProductId

	valid := validation.Validation{}
	valid.Required(price, "price").Message("单价不能为空!")
	valid.Required(count, "count").Message("数量不能为空")
	valid.Required(productId, "productId").Message("产品ID不能为空")
	
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if (models.ExistProductId(id)) {
			data := make(map[string]interface{})
			if count != 0 {
				data["count"] = count
			}
			if price != 0 {
				data["price"] = price
			}
			if url != "" {
				data["url"] = url
			}
			if productId != 0 {
				data["productId"] = productId
			}
			models.EditRecords(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_PRODUCT
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}
func DeleteRecords (ctx * gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()

	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("id必须大于0")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistRecordId(id) {
			models.DeleteRecords(id)
		} else {
			code = e.ERROR_NOT_EXIST_PRODUCT
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}