package routes

import (
	"kwc-backend/controllers"
	"kwc-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func ContestantRoutes(app *gin.Engine) {
	app.POST("/api/contestants", middlewares.Auth([]string{"admin"}), controllers.AddContestant)
	app.GET("/api/contestants", controllers.ListContestants)
	app.GET("/api/contestants/:id", controllers.GetContestant)
	app.PATCH("/api/contestants/:id", middlewares.Auth([]string{"admin"}), controllers.UpdateContestant)
	app.DELETE("/api/contestants/:id", middlewares.Auth([]string{"admin"}), controllers.DeleteContestant)
}
