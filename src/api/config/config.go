package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	secretGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	LogLevel                = "info"
)

var apiGithubAccessToken string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiGithubAccessToken = os.Getenv("TOKEN_GITHUB")
}

func GetGithubAccessToken() string {
	return apiGithubAccessToken
}

func IsProduction() bool {
	return os.Getenv("GO_ENVIRONMENT") == "production"
}