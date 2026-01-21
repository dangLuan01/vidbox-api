package v1handler

import (
	"net/http"

	v1service "vidbox-api/internal/service/v1"

	"github.com/google/uuid"

	"vidbox-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service v1service.UserService
}
type GetUserByUUIDParam struct{
	Uuid uuid.UUID `uri:"uuid" binding:"uuid"`
}
func NewUserHandler(service v1service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
//OPHIM
func (uh *UserHandler) Crawler(ctx *gin.Context) {
	if err := uh.service.Crawler(); err != nil {
		utils.ResponseError(ctx, err)
	}

	utils.ResponseSatus(ctx, http.StatusOK)

}

//KKPHIM
func (uh *UserHandler) CrawlerTvKkphim(ctx *gin.Context) {
	if err := uh.service.CrawlerTvKkphim(); err != nil {
		utils.ResponseError(ctx, err)
	}

	utils.ResponseSatus(ctx, http.StatusOK)
}

func (uh *UserHandler) CrawlerMovieKkphim(ctx *gin.Context) {
	if err := uh.service.CrawlerMovieKkphim(); err != nil {
		utils.ResponseError(ctx, err)
	}

	utils.ResponseSatus(ctx, http.StatusOK)
}

func (uh *UserHandler) CrawlerAllKkphim(ctx *gin.Context) {
	if err := uh.service.CrawlerAllKKphim(); err != nil {
		utils.ResponseError(ctx, err)
	}

	utils.ResponseSatus(ctx, http.StatusOK)
}
//NGUONC
func (uh *UserHandler) CrawlerTvNguonC(ctx *gin.Context) {
	if err := uh.service.CrawlerTvNguonC(); err != nil {
		utils.ResponseError(ctx, err)
	}

	utils.ResponseSatus(ctx, http.StatusOK)
}
