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
	if req.Email != nil && req.Password == nil || req.Email == nil && req.Password != nil {
		common.AbortWithError(c, common.ErrInvalidEmailOrPassWord)
		return
	}
	if req.AccessToken != nil && req.IdToken == nil || req.AccessToken == nil && req.IdToken != nil {
		common.AbortWithError(c, common.ErrInvalidToken)
		return
	}

	token, err := h.service.LoginUser(c, req)
	if err != nil {
		common.AbortWithError(c, err)
		return
	}

	c.JSON(common.SUCCESS_STATUS, token)
}
