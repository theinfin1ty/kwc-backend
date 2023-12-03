package routes

import (
	"kwc-backend/controllers"

	"github.com/gin-gonic/gin"
)

func EpisodeRoutes(app *gin.Engine) {
	app.POST("/api/episodes", controllers.CreateEpisode)
	app.GET("/api/episodes", controllers.ListEpisodes)
	app.GET("/api/episodes/:id", controllers.GetEpisode)
	app.GET("/api/episodes/season/:id", controllers.ListEpisodesBySeason)
	app.PATCH("/api/episodes/:id", controllers.UpdateEpisode)
	app.DELETE("/api/episodes/:id", controllers.DeleteEpisode)
}