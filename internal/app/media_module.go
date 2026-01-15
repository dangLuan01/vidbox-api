package app

import (
	v1handler "vidbox-api/internal/handler/v1"
	"vidbox-api/internal/repository"
	"vidbox-api/internal/routes"
	v1routes "vidbox-api/internal/routes/v1"
	v1service "vidbox-api/internal/service/v1"
)

type MediaModule struct {
	routes routes.Route
}

func NewMediaModule(ctx *ModuleContext) *MediaModule {

	mediaRepo := repository.NewSqlMediaRepository(ctx.DB)
	mediaService := v1service.NewMediaService(mediaRepo)
	mediaHandler := v1handler.NewMediaHandler(mediaService)
	mediaRoutes := v1routes.NewMediaRoutes(mediaHandler)

	return &MediaModule{
		routes: mediaRoutes,
	}
}
func (m *MediaModule) Routes() routes.Route {
	return m.routes
}