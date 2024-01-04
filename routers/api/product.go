package api

import (
	"fmt"
	"net/http"

	"github.com/Leexiaop/molars_rd/models"
	"github.com/Leexiaop/molars_rd/pkg/e"
	"github.com/Leexiaop/molars_rd/pkg/setting"
	"github.com/Leexiaop/molars_rd/pkg/util"
	"github.com/gin-gonic/gin"

	"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"
)

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
	name := ctx.PostForm("name")
	price := com.StrTo(ctx.Query("price")).MustInt()
	url := ctx.PostForm("url")
	fmt.Print(name, price, url)


	valid := validation.Validation{}
	valid.Required(name, "name").Message("产品名称不能为空")
	valid.Required(price, "price").Message("单价不能为空")
	valid.Required(url, "url").Message("图片不能为空")
	
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		fmt.Print(444)
		if !models.ExistProductName(name) {
			code = e.SUCCESS
			models.AddProducts(name, price, url)
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