package routes

import (
	"kwc-backend/controllers"
	"kwc-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(app *gin.Engine) {
	app.POST("/api/users", middlewares.Auth([]string{"admin"}), controllers.CreateUser)
	app.POST("/api/token", controllers.GetUserToken)
	app.GET("/api/users", middlewares.Auth([]string{"admin"}), controllers.ListUsers)
	app.GET("/api/users/:id", controllers.GetUser)
	app.PATCH("/api/users/:id", controllers.UpdateUser)
	app.DELETE("/api/users/:id", controllers.DeleteUser)
}
