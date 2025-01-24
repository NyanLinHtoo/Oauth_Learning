package handler

import (
	"log"
	"net/http"

	"github.com/NyanLinHtoo/Oauth_Learning/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc services.AuthServiceInterface
}

func NewAuthHandler(svc services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

// Google Login Handler
func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	authURL := h.svc.GetGoogleAuthURL("state-token")
	log.Println("C Redirect ", authURL)
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

// Google Callback Handler
func (h *AuthHandler) GoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != "state-token" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
		return
	}

	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code parameter"})
		return
	}

	token, err := h.svc.ExchangeCodeForToken(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
