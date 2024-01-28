package routers

import (
	"net/http"

	_ "github.com/Leexiaop/molars_rd/docs"
	jwt "github.com/Leexiaop/molars_rd/middleware"
	"github.com/Leexiaop/molars_rd/pkg/export"
	"github.com/Leexiaop/molars_rd/pkg/upload"
	"github.com/Leexiaop/molars_rd/routers/api"
	"github.com/Leexiaop/molars_rd/routers/api/basic"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() * gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/api/auth", api.GetAuth)
	r.POST("api/upload", api.Upload)
	r.StaticFS("/api/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/api/upload/images", http.Dir(upload.GetImagesFullPath()))
	basicApi := r.Group("/api/basic")
	basicApi.Use(jwt.JWT())
	{
		//	商品的增删改查
		basicApi.POST("/product", basic.AddProducts)
		basicApi.DELETE("/product/:id", basic.DeleteProducts)
		basicApi.PUT("/product", basic.EditProducts)
		basicApi.GET("/product", basic.GetProducts)

		//	记录的增删改查
		basicApi.GET("/record", basic.GetRecords)
		basicApi.POST("/record", basic.AddRecords)
		basicApi.PUT("/record", basic.EditRecords)
		basicApi.DELETE("/record/:id", basic.DeleteRecords)
		basicApi.GET("/record/export", basic.ExportRecords)

		//	用户增删改查
		basicApi.GET("/user", basic.GetUsers)
	}
	return r
}