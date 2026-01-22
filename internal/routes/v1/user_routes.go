package v1routes

import (
	v1handler "vidbox-api/internal/handler/v1"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	handler *v1handler.UserHandler
}

func NewUserRoutes(handler *v1handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		handler: handler,
	}
}

func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	// users := r.Group("/users")
	// {
	// 	users.GET("", ur.handler.GetAllUser)
	// 	users.GET("/:uuid", ur.handler.GetUserByUUID)
	// 	users.POST("", ur.handler.CreateUser)
	// 	users.PUT("/:uuid", ur.handler.UpdateUser)
	// 	users.DELETE("/:uuid", ur.handler.DeleteUser)
	// }
	users := r.Group("/crawler") 
	{
		//ophim
		users.GET("", ur.handler.Crawler)
		// users.GET("/kkphim/tv", ur.handler.CrawlerTvKkphim)
		// users.GET("/kkphim/movie", ur.handler.CrawlerMovieKkphim)
		//kkphim
		users.GET("/kkphim/all", ur.handler.CrawlerAllKkphim)
		//nguonc
		users.GET("/nguonc/all", ur.handler.CrawlerAllNguonC)		
	}
}
