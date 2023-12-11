package routes

import (
	"kwc-backend/controllers"
	"kwc-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func QuestionRoutes(app *gin.Engine) {
	app.POST("/api/questions", middlewares.Auth([]string{"admin"}), controllers.CreateQuestion)
	app.GET("/api/questions/:id", controllers.ListQuestions)
	app.PATCH("/api/questions/:id", middlewares.Auth([]string{"admin"}), controllers.UpdateQuestion)
	app.DELETE("/api/questions/:id", middlewares.Auth([]string{"admin"}), controllers.DeleteQuestion)
}
