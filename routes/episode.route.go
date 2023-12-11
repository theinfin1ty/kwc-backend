package routes

import (
	"kwc-backend/controllers"
	"kwc-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func EpisodeRoutes(app *gin.Engine) {
	app.POST("/api/episodes", middlewares.Auth([]string{"admin"}), controllers.CreateEpisode)
	app.GET("/api/episodes/:id", controllers.ListEpisodesBySeason)
	app.PATCH("/api/episodes/:id", middlewares.Auth([]string{"admin"}), controllers.UpdateEpisode)
	app.DELETE("/api/episodes/:id", middlewares.Auth([]string{"admin"}), controllers.DeleteEpisode)
}
