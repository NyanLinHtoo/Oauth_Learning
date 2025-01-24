package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Oauth2Config struct {
	// Oauth2 Config
	OauthClientID     string
	OauthClientSecret string
	RedirectURL       string
	Scopes            []string
	Endpoint          oauth2.Endpoint
}

var GoogleOauthConfig *oauth2.Config

func LoadConfig() {
	// Load .env file
	err := godotenv.Load()
	log.Println("Got Error ===>", err)
	if err != nil {
		log.Println("No .env file found, loading environment variables")
	}

	GoogleOauthConfig = &oauth2.Config{
		ClientID:     getEnv("OAUTH2_ClientID", "1072416741776-0u4931m6io4q69iqhlkh8q40hov5eump.apps.googleusercontent.com"),
		ClientSecret: getEnv("OAUTH2_ClientSecret", "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"),
		RedirectURL:  getEnv("OAUTH2_RedirectURL", "http://localhost:8080/api/auth/google/callback"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
