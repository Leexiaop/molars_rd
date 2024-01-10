package routers

import (
	"github.com/Leexiaop/molars_rd/pkg/setting"
	"github.com/Leexiaop/molars_rd/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() * gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	route := r.Group("/api")
	{
		//	商品的增删改查
		route.POST("/product", api.AddProducts)
		route.DELETE("/product/:id", api.DeleteProducts)
		route.PUT("/product", api.EditProducts)
		route.GET("/product", api.GetProducts)

		//	记录的增删改查
		route.GET("/record", api.GetRecords)
		route.POST("/record", api.AddRecords)
		route.PUT("/record", api.EditRecords)
		route.DELETE("/record/:id", api.DeleteRecords)
	}
	return r
}