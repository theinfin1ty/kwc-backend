package routes

import (
	"kwc-backend/controllers"
	"kwc-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SeasonRoutes(app *gin.Engine) {
	app.POST("/api/seasons", middlewares.Auth([]string{"admin"}), controllers.CreateSeason)
	app.GET("/api/seasons", controllers.ListSeasons)
	app.GET("/api/seasons/:id", controllers.GetSeason)
	app.PATCH("/api/seasons/:id", middlewares.Auth([]string{"admin"}), controllers.UpdateSeason)
	app.DELETE("/api/seasons/:id", middlewares.Auth([]string{"admin"}), controllers.DeleteSeason)
}
