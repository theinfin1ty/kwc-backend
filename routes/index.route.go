package routes

import "github.com/gin-gonic/gin"

func Routes(app *gin.Engine) {
	UserRoutes(app)
	ContestantRoutes(app)
	SeasonRoutes(app)
	EpisodeRoutes(app)
	QuestionRoutes(app)
}
