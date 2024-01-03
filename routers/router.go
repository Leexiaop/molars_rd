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
		route.GET("/product", api.GetProducts)
	}
	return r
}