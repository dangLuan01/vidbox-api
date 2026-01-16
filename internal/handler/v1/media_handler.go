package v1handler

import (
	"net/http"
	v1dto "vidbox-api/internal/dto/v1"
	v1service "vidbox-api/internal/service/v1"
	"vidbox-api/internal/utils"
	"vidbox-api/internal/validation"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	service v1service.MediaService
}

func NewMediaHandler(service v1service.MediaService) *MediaHandler {
	return &MediaHandler{
		service: service,
	}
}

func (mh *MediaHandler) GetMediaTv(ctx *gin.Context) {
	var param v1dto.MediaInput

	err := ctx.ShouldBindUri(&param)
	if err != nil {

		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return 
	}
	param.MediaType = "tv"
	slug, err := mh.service.GetMedia(param)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.",slug)
}

func (mh *MediaHandler) GetMediaMovie(ctx *gin.Context) {
	var param v1dto.MediaInput

	err := ctx.ShouldBindUri(&param)
	if err != nil {

		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return 
	}
	param.MediaType = "movie"
	slug, err := mh.service.GetMedia(param)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.",slug)
}