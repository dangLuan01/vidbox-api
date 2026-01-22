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
func (uh *UserHandler) CrawlerAllKkphim(ctx *gin.Context) {
	if err := uh.service.CrawlerAllKKphim(); err != nil {
		utils.ResponseError(ctx, err)
	}

	utils.ResponseSatus(ctx, http.StatusOK)
}

//NGUONC
func (uh *UserHandler) CrawlerAllNguonC(ctx *gin.Context) {

	go uh.service.CrawlerAllNguonC()
	
	utils.ResponseSatus(ctx, http.StatusOK)
}
