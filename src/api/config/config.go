package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var apiGithubAccessToken string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiGithubAccessToken = os.Getenv("TOKEN_GITHUB")

	fmt.Println("======== apiGithubAccessToken ========", apiGithubAccessToken)
}

// Getter
func GetGithubAccessToken() string {
	return apiGithubAccessToken
}