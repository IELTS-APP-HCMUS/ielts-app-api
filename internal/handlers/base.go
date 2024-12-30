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

func (h *Handler) RegisterRoutes(c *gin.Engine) {

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
	authRoutes := c.Group("/api/auth")
	{
		authRoutes.POST("/request-reset-password", h.RequestResetPassword)
		authRoutes.POST("/validate-otp", h.ValidateOTP)
		authRoutes.POST("/reset-password", h.ResetPassword)
	}

	quizzes := c.Group("/v1/quizzes")
	{
		quizzes.GET("/:quiz_id", middleware.UserAuthentication, h.GetQuiz())
		quizzes.GET("", middleware.OptionalUserAuthentication(), h.GetQuizzes())
		quizzes.POST("/:quiz_id/answer", middleware.UserAuthentication, h.SubmitQuiz())
	}

	tagSearches := c.Group("/v1/tag-searches")
	{
		tagSearches.GET("", h.GetTagSearches())
	}

	answerRoutes := c.Group("/v1/answers")
	{
		answerRoutes.GET("/:answer_id", middleware.UserAuthentication, h.GetAnswer)
		answerRoutes.GET("/statistics", middleware.UserAuthentication, h.GetAnswerStatistic)
	}

	vocab := c.Group("/v1/vocabs")
	{
		vocab.GET("", middleware.UserAuthentication, h.GetVocab)
		vocab.POST("", middleware.UserAuthentication, h.CreateVocab)
		vocab.PATCH("", middleware.UserAuthentication, h.UpdateVocab)
		vocab.DELETE("", middleware.UserAuthentication, h.DeleteVocab)
		vocab.GET("/reading", middleware.UserAuthentication, h.GetReadingVocab())
	}

	plan := c.Group("/v1/plans")
	{
		plan.GET("", middleware.UserAuthentication, h.GetPlan)
		plan.POST("", middleware.UserAuthentication, h.CreatePlan)
	}

	masterDataRoutes := c.Group("/api/v1/master-data")
	{
		masterDataRoutes.GET("", h.GetMasterData)
	}

}
