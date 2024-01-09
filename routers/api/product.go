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

type BodyStruct struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Url string `json:"url"`
}

//	获取产品列表

func GetProducts (ctx * gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS

	data["list"] = models.GetProducts(util.GetPage(ctx), setting.PageSize, maps)
	data["total"] = models.GetProductsTotal(maps)

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}

func AddProducts (ctx * gin.Context) {
	jsonData, _ := ctx.GetRawData()

	var m BodyStruct

	json.Unmarshal(jsonData, &m)
	name := m.Name
	price := m.Price
	url := m.Url

	createdBy := ctx.DefaultQuery("createdBy", "13691388204")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("产品名称不能为空")
	valid.Required(price, "price").Message("单价不能为空")
	valid.Required(url, "url").Message("图片不能为空")
	
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistProductName(name) {
			code = e.SUCCESS
			models.AddProducts(name, price, url, createdBy)
		} else {
			code = e.ERROR_EXIST_PRODUCT
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditProducts (ctx * gin.Context) {
	jsonData, _ := ctx.GetRawData()

	var m BodyStruct
	json.Unmarshal(jsonData, &m)
	id := m.ID
	name := m.Name
	price := m.Price
	url := m.Url
	modifiedBy := ctx.DefaultQuery("modifiedBy", "13691388204")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("找不到对应的ID")
	valid.Required(name, "name").Message("产品名称不能为空")
	valid.Required(price, "price").Message("单价不能为空")
	valid.Required(url, "url").Message("图片不能为空")
	
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if (models.ExistProductId(id)) {
			data := make(map[string]interface{})
			data["modifield_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if price != 0 {
				data["price"] = price
			}
			if url != "" {
				data["url"] = url
			}
			models.EditProducts(id, data)
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

func DeleteProducts (ctx * gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()

	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("id必须大于0")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistProductId(id) {
			models.DeleteProducts(id)
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