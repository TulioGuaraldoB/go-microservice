package env

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func GetEnvironmentVariables() {
	if err := godotenv.Load(".env"); err != nil {
		errMessage := fmt.Sprintf("Failed to read .env file. %s", err.Error())
		log.Fatal(errMessage)
	}
}
