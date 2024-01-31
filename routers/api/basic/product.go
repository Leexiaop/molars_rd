package basic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Leexiaop/molars_rd/models"
	"github.com/Leexiaop/molars_rd/pkg/app"
	"github.com/Leexiaop/molars_rd/pkg/e"
	"github.com/Leexiaop/molars_rd/pkg/setting"
	"github.com/Leexiaop/molars_rd/pkg/util"
	"github.com/Leexiaop/molars_rd/service/product_service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/astaxie/beego/validation"
)

func GetProducts (ctx * gin.Context) {
	appG := app.Gin{C: ctx}

	pageSize, _ :=strconv.Atoi(ctx.DefaultQuery("pageSize", fmt.Sprint(setting.AppSetting.PageSize)))
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", fmt.Sprint(util.GetPage(ctx))))

	productService := product_service.Product{
		PageNum: pageNum,
		PageSize: pageSize,
	}

	products, err := productService.GetAll()


	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_PRODUCTS_FAIL, nil)
		return
	}

	count, err := productService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_PRODUCTS_FAIL, nil)
		return
	
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list": products,
		"pageSize": pageSize,
		"pageNum": pageNum,
		"total": count,
	})
}

func AddProducts (ctx * gin.Context) {
	appG := app.Gin{C: ctx}
	username, _ := ctx.Get("username")
	jsonData, _ := ctx.GetRawData()

	var m models.Product

	json.Unmarshal(jsonData, &m)
	name := m.Name
	price := m.Price
	url := m.Url
	created_by := username.(string)

	valid := validation.Validation{}
	valid.Required(name, "name").Message("产品名称不能为空")
	valid.Required(price, "price").Message("单价不能为空")
	valid.Required(url, "url").Message("图片不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	productService := product_service.Product{
		Name: name,
		Price: price,
		Url: url,
		CreatedBy: created_by,
	}
	exists, err := productService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_PRODUCT_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_PRODUCT, nil)
		return
	}

	err = productService.Add()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_PRODUCT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func EditProducts (ctx * gin.Context) {
	appG := app.Gin{C: ctx}
	username, _ := ctx.Get("username")

	jsonData, _ := ctx.GetRawData()

	var m models.Product
	json.Unmarshal(jsonData, &m)
	id := m.ID
	name := m.Name
	price := m.Price
	url := m.Url
	modifieldBy := username.(string)

	valid := validation.Validation{}
	valid.Required(id, "id").Message("找不到对应的ID")
	valid.Required(name, "name").Message("产品名称不能为空")
	valid.Required(price, "price").Message("单价不能为空")
	valid.Required(url, "url").Message("图片不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	productService := product_service.Product{
		ID: id,
		Name: name,
		Price: price,
		Url: url,
		ModifieldBy: modifieldBy,
	}
	exists, err := productService.ExistById()

	 if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_PRODUCT_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_PRODUCT, nil)
		return
	}
	err = productService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_PRODUCT_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteProducts (ctx * gin.Context) {
	appG := app.Gin{C: ctx}
	valid := validation.Validation{}
	id := com.StrTo(ctx.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	productService := product_service.Product{ID: id}
	exists, err := productService.ExistById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_PRODUCT, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_PRODUCT, nil)
		return
	}

	if err := productService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_PRODUCT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}