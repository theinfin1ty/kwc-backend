package routes

import (
	"kwc-backend/controllers"

	"github.com/gin-gonic/gin"
)

func QuestionRoutes(app *gin.Engine) {
	app.POST("/api/questions", controllers.CreateQuestion)
	app.GET("/api/questions/:id", controllers.ListQuestions)
	app.PATCH("/api/questions/:id", controllers.UpdateQuestion)
	app.DELETE("/api/questions/:id", controllers.DeleteQuestion)
}
