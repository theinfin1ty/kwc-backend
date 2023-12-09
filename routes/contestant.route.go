package routes

import (
	"kwc-backend/controllers"
	"kwc-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func ContestantRoutes(app *gin.Engine) {
	app.POST("/api/contestant", middlewares.Auth([]string{"admin"}), controllers.AddContestant)
	app.GET("/api/contestant", controllers.ListContestants)
	app.GET("/api/contestant/:id", controllers.GetContestant)
	app.PATCH("/api/contestant/:id", middlewares.Auth([]string{"admin"}), controllers.UpdateContestant)
	app.DELETE("/api/contestant/:id", middlewares.Auth([]string{"admin"}), controllers.DeleteContestant)
}
