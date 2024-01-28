package api

import (
	"net/http"

	"github.com/Leexiaop/molars_rd/pkg/app"
	"github.com/Leexiaop/molars_rd/pkg/e"
	"github.com/Leexiaop/molars_rd/pkg/logging"
	"github.com/Leexiaop/molars_rd/pkg/upload"
	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	appG := app.Gin{C: ctx}

	file, image, err := ctx.Request.FormFile("file")
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImagesFullPath()
	savePath := upload.GetImagesPath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := ctx.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url": upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}