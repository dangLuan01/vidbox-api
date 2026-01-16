package app

import (
	"log"

	"vidbox-api/internal/config"
	"vidbox-api/internal/db"
	"vidbox-api/internal/routes"
	"vidbox-api/internal/validation"

	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	config *config.Config
	router *gin.Engine
	modules []Module
}

type ModuleContext struct {
	DB *goqu.Database
}

func NewApplication(cfg *config.Config) (*Application, error) {

	if err := validation.InitValidator(); err != nil {
		log.Fatalf("⛔ Validation init failed %v:", err)
		return nil, err
	}
	
	r := gin.Default()
	
	if err := db.InitDB(); err != nil {
		log.Fatalf("⛔ Unable to connect to sql")
		return nil, err
	}

	ctx := &ModuleContext{
		DB: db.DB,
	}

	modules := []Module{
		NewUserModule(ctx),
		NewMediaModule(ctx),
	}

	routes.RegisterRoute(r,getModuleRoutes(modules)...)

	return &Application{
		config: cfg,
		router: r,
		modules: modules,
	}, nil
}

func (a *Application) Run() error {
	
	return a.router.Run(a.config.ServerAddress)
}

func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))
	for i, module := range modules {
		routeList[i] = module.Routes()
	}

	return routeList
}
func LoadEnv()  {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}