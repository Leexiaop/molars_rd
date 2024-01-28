package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Leexiaop/molars_rd/models"
	"github.com/Leexiaop/molars_rd/pkg/gredis"
	"github.com/Leexiaop/molars_rd/pkg/logging"
	"github.com/Leexiaop/molars_rd/pkg/setting"
	"github.com/Leexiaop/molars_rd/pkg/util"
	"github.com/Leexiaop/molars_rd/routers"
)
func init () {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

func main () {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
