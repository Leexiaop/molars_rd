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
		route.POST("/product", api.AddProducts)
		route.DELETE("/product/:id", api.DeleteProducts)
		route.PUT("/product", api.EditProducts)
		route.GET("/product", api.GetProducts)
	}
	{
		route.GET("/record", api.GetRecords)
	}
	return r
}