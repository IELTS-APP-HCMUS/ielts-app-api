package handlers

import (
	services "ielts-app-api/internal/services"
	"ielts-app-api/middleware"

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

// Define API route here
func (h *Handler) RegisterRoutes(c *gin.Engine) {
	// Health check
	health := c.Group("api/health")
	{
		health.GET("/status", h.CheckStatusHealth)
	}
	userRoutes := c.Group("/api/users")
	{
		userRoutes.GET("", middleware.UserAuthentication, h.GetUserProfile)
		userRoutes.POST("/signup", h.SignUp)
		userRoutes.POST("/login", h.LogIn)
		userRoutes.GET("/target", middleware.UserAuthentication, h.GetTarget)
		userRoutes.POST("/target", middleware.UserAuthentication, h.CreateTarget)
		userRoutes.PATCH("/target", middleware.UserAuthentication, h.UpdateTarget)
	}

	quizzes := c.Group("/v1/quizzes")
	{
		// API Get Quiz Detail
		// quizzes.GET("/:quiz_id", middleware.OptionalUserAuthentication(), h.GetQuiz())

		// API Get Quiz Answer
		// quizzes.GET("/answers/:answer_id", middleware.UserAuthentication(), h.GetQuizAnswer())

		//API Listing Quiz
		quizzes.GET("", middleware.OptionalUserAuthentication(), h.GetQuizzes())
	}
}
