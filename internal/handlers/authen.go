package handlers

import (
	"ielts-app-api/common"
	"ielts-app-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var req models.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.AbortWithError(c, common.ErrInvalidInput)
		return
	}

	err := h.service.SignupUser(c, req)
	if err != nil {
		common.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, common.ResponseOk("User created successfully"))
}

func (h *Handler) LogIn(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.AbortWithError(c, common.ErrInvalidInput)
		return
	}

	if req.Email == nil && req.IdToken == nil {
		common.AbortWithError(c, common.ErrInvalidInput)
		return
	}

	token, err := h.service.LoginUser(c, req)
	if err != nil {
		common.AbortWithError(c, err)
		return
	}

	c.JSON(common.SUCCESS_STATUS, common.BaseResponseMess(common.SUCCESS_STATUS, "Login successfully", token))
}

func (h *Handler) RequestResetPassword(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	otp, err := h.service.GenerateOTP(c.Request.Context(), req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
		return
	}

	err = common.SendOTPEmail(common.FromEmail, req.Email, otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

func (h *Handler) ValidateOTP(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
		OTP   string `json:"otp" binding:"required"`
	}

	// Bind and validate the request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Call the service to validate the OTP
	err := h.service.ValidateOTP(c.Request.Context(), req.Email, req.OTP)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP validated successfully"})
}

func (h *Handler) ResetPassword(c *gin.Context) {
	var req models.ResetPasswordRequest

	// Validate the incoming request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Call the service to reset the password
	err := h.service.ResetPassword(c.Request.Context(), req.Email, req.NewPassword)
	if err != nil {
		if err.Error() == "email not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset password"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}
