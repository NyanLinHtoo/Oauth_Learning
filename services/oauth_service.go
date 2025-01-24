package services

import (
	"context"
	"log"

	"github.com/NyanLinHtoo/Oauth_Learning/config"
	"golang.org/x/oauth2"
)

type AuthServiceInterface interface {
	GetGoogleAuthURL(state string) string
	ExchangeCodeForToken(ctx context.Context, code string) (*oauth2.Token, error)
}

type AuthService struct{}

func NewAuthService() AuthServiceInterface {
	return &AuthService{}
}

func (s *AuthService) GetGoogleAuthURL(state string) string {
	log.Println("Client Id=============>", config.GoogleOauthConfig.ClientID)
	authURL := config.GoogleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	log.Printf("Generated Auth URL:==============> %s ", authURL)
	return authURL
}

func (s *AuthService) ExchangeCodeForToken(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := config.GoogleOauthConfig.Exchange(ctx, code)
	if err != nil {
		log.Printf("Error exchanging code for token: %s", err)
		return nil, err
	}
	return token, nil
}
