package basic

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Leexiaop/molars_rd/models"
	"github.com/Leexiaop/molars_rd/pkg/app"
	"github.com/Leexiaop/molars_rd/pkg/e"
	"github.com/Leexiaop/molars_rd/pkg/export"
	"github.com/Leexiaop/molars_rd/pkg/setting"
	"github.com/Leexiaop/molars_rd/pkg/util"
	"github.com/Leexiaop/molars_rd/service/product_service"
	"github.com/Leexiaop/molars_rd/service/record_service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/astaxie/beego/validation"
)

//	获取记录列表
func GetRecords (ctx * gin.Context) {
	appG := app.Gin{C: ctx}
	productId, _ := strconv.Atoi(ctx.Query("product_id"))

	valid := validation.Validation{}

	if (valid.HasErrors()) {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PRODUCT_PARAMS, nil)
		return
	}
	recordServeice := record_service.Record{
		ProductId: productId,
		PageNum: util.GetPage(ctx),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := recordServeice.Counts()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_RECORDS_FAIL, nil)
		return
	}

	records, err := recordServeice.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list": records,
		"total": total,
	})
}

func AddRecords (ctx * gin.Context) {
	appG := app.Gin{C: ctx}
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

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	productService := product_service.Product{
		ID: productId,
	}
	exists, err := productService.ExistById()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_PRODUCT_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_PRODUCT, nil)
		return
	}

	recordService := record_service.Record{
		ProductId: productId,
		Price: price,
		Count: count,
		Url: url,
	}

	if err := recordService.Add();err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_RECORD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
func EditRecords (ctx * gin.Context) {
	appG := app.Gin{C: ctx}
	jsonData, _ := ctx.GetRawData()

	var m models.Record
	json.Unmarshal(jsonData, &m)
	id := m.ID
	price := m.Price
	count := m.Count
	url := m.Url
	productId := m.ProductId

	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不能为空!")
	valid.Required(price, "price").Message("单价不能为空!")
	valid.Required(count, "count").Message("数量不能为空")
	valid.Required(productId, "productId").Message("产品ID不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	recordService := record_service.Record{
		ID: id,
		Price: price,
		Count: count,
		Url: url,
		ProductId: productId,
	}

	exists, err := recordService.ExistById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_RECORD_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_RECORD, nil)
		return
	}

	productSerice := product_service.Product{
		ID: productId,
	}
	exists, err = productSerice.ExistById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_PRODUCT_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_PRODUCT, nil)
		return
	}

	err = recordService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_RECORD_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
func DeleteRecords (ctx * gin.Context) {
	appG := app.Gin{C: ctx}
	valid := validation.Validation{}
	id := com.StrTo(ctx.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	recordService := record_service.Record{
		ID: id,
	}

	exits, err := recordService.ExistById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_RECORD_FAIL, nil)
		return
	}
	if !exits {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_RECORD, nil)
		return
	}
	if err := recordService.Delete();err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_RECORD_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func ExportRecords(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	productId, _ := strconv.Atoi(ctx.Query("productId"))

	recordService := record_service.Record{
		ProductId: productId,
	}

	filename, err := recordService.Export()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXPORT_RECORD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"export_url": export.GetExcelFullUrl(filename),
		"export_save_url": export.GetExcelPath() + filename,
	})
}