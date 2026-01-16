package v1routes

import (
	v1handler "vidbox-api/internal/handler/v1"

	"github.com/gin-gonic/gin"
)

type MediaRoutes struct {
	handler *v1handler.MediaHandler
}

func NewMediaRoutes(handler *v1handler.MediaHandler) *MediaRoutes {
	return &MediaRoutes{
		handler: handler,
	}
}

func (mr *MediaRoutes) Register(r *gin.RouterGroup) {
	
	medias := r.Group("/")
	{
		medias.GET("tv/:tmdb_id/:season", mr.handler.GetMediaTv)
		medias.GET("movie/:tmdb_id", mr.handler.GetMediaMovie)
	}
}