package routes

import (
	"kwc-backend/controllers"

	"github.com/gin-gonic/gin"
)

func ContestantRoutes(app *gin.Engine) {
	app.POST("/api/contestants", controllers.AddContestant)
	app.GET("/api/contestants", controllers.ListContestants)
	app.GET("/api/contestants/:id", controllers.GetContestant)
	app.PATCH("/api/contestants/:id", controllers.UpdateContestant)
	app.DELETE("/api/contestants/:id", controllers.DeleteContestant)
}
