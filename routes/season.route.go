package routes

import (
	"kwc-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SeasonRoutes(app *gin.Engine) {
	app.POST("/api/seasons", controllers.CreateSeason)
	app.GET("/api/seasons", controllers.ListSeasons)
	app.GET("/api/seasons/:id", controllers.GetSeason)
	app.PATCH("/api/seasons/:id", controllers.UpdateSeason)
	app.DELETE("/api/seasons/:id", controllers.DeleteSeason)
}
