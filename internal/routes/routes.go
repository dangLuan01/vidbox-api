package routes

import (
	"vidbox-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoute(r *gin.Engine, routes ...Route) {
	v1api := r.Group("/api/v1")
	v1api.Use(	
		//middleware.ApiKeyMiddleware(),
		middleware.CORSMiddleware(),
		middleware.RateLimiterMiddleware(), 
	)
	
	for _, route := range routes {
		route.Register(v1api)
	}

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"error":"NOT FOUND",
			"path": ctx.Request.URL.Path,
		})
	})
}