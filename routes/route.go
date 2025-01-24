package routes

import (
	"github.com/NyanLinHtoo/Oauth_Learning/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	AuthHandler *handler.AuthHandler
}

func NewRouter(auth *handler.AuthHandler) *Router {
	return &Router{
		AuthHandler: auth,
	}
}

func (ro *Router) RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{

		// Auth routes
		authApi := apiGroup.Group("/auth")
		authApi.GET("/google/login", ro.AuthHandler.GoogleLogin)
		authApi.GET("/google/callback", ro.AuthHandler.GoogleCallback)
	}
}
