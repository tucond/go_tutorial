package routes

import (
	"github.com/labstack/echo"
	"youtube-manager-go/middlewares"
	"youtube-manager-go/web/api"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
		g.GET("/video/:id", api.GetVideo())
		g.GET("/related/:id", api.FetchRelatedVideos())
		g.GET("/search", api.SearchVideos())
	}

	fg := g.Group("/favorite", middlewares.FirebaseGuard())
	{
		fg.POST("/:id/toggle", api.ToggleFavoriteVideo())
	}
}
