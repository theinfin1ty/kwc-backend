package routes

import (
	"kwc-backend/controllers"
	"kwc-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(app *gin.Engine) {
	app.POST("/api/users", middlewares.Auth([]string{"admin"}), controllers.CreateUser)
	// app.POST("/api/token", controllers.GetUserToken)
	// app.GET("/api/users", middlewares.Auth([]string{"admin"}), controllers.ListUsers)
	app.GET("/api/users/:id", middlewares.Auth([]string{"admin"}), controllers.GetUser)
	app.PATCH("/api/users/:id", middlewares.Auth([]string{"admin"}), controllers.UpdateUser)
	app.DELETE("/api/users/:id", middlewares.Auth([]string{"admin"}), controllers.DeleteUser)
}
