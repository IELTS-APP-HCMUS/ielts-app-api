package handlers

import (
	services "ielts-app-api/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(c *gin.Engine) {
	userRoutes := c.Group("/api/users")
	{
		userRoutes.POST("/signup", h.SignUp)
		userRoutes.POST("/login", h.LogIn)
		userRoutes.POST("/logout", h.LogOut)
	}
}
