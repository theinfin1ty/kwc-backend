package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {
	_ = godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// }
	return os.Getenv(key)
}
