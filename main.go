package main

import (
	"fmt"
	"kwc-backend/configs"
	"kwc-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if configs.GetEnvVariable("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()

	app.Use(cors.Default())

	configs.ConnectDB()
	configs.InitializeFirebase()

	routes.Routes(app)

	fmt.Printf("Server is running on port %s \n", configs.GetEnvVariable("PORT"))

	app.Run(":" + configs.GetEnvVariable("PORT"))
}
