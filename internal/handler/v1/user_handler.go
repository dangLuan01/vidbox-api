package v1handler

import (
	"net/http"

	v1service "vidbox-api/internal/service/v1"
	"vidbox-api/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service v1service.UserService
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

	go uh.service.CrawlerTvNguonC()
	
	utils.ResponseSatus(ctx, http.StatusOK)
}

func (uh *UserHandler) CrawlerMovieNguonC(ctx *gin.Context) {

	go uh.service.CrawlerMovieNguonC()
	
	utils.ResponseSatus(ctx, http.StatusOK)
}
