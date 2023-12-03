package main

import (
	"fmt"
	"kwc-backend/configs"
	"kwc-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	configs.ConnectDB()
	configs.InitializeFirebase()

	routes.Routes(app)

	fmt.Printf("Server is running on port %s \n", "3000")

	app.Run(":" + "3000")
}
