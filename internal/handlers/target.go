package handlers

import (
	"ielts-app-api/common"
	"ielts-app-api/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTarget(c *gin.Context) {
	ok, userJWTProfile := common.ProfileFromJwt(c)
	if !ok {
		c.JSON(common.INTERNAL_SERVER_ERR, gin.H{"error": "User not found"})
		return
	}
	data, err := h.service.GetTargetById(c, userJWTProfile.Id)
	if err != nil {
		c.JSON(common.INTERNAL_SERVER_ERR, gin.H{"error": err.Error()})
		return
	}
	c.JSON(common.SUCCESS_STATUS, gin.H{"message": "Get user succerfully", "data": data})
}

func (h *Handler) CreateTarget(c *gin.Context) {
	ok, userJWTProfile := common.ProfileFromJwt(c)
	if !ok {
		c.JSON(common.INTERNAL_SERVER_ERR, gin.H{"error": "User not found"})
		return
	}
	var req models.TargetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(common.INTERNAL_SERVER_ERR, gin.H{"error": err.Error()})
		return
	}
	data, err := h.service.CreateTarget(c, userJWTProfile.Id, req)
	if err != nil {
		c.JSON(common.INTERNAL_SERVER_ERR, gin.H{"error": err.Error()})
		return
	}
	c.JSON(common.SUCCESS_STATUS, gin.H{"message": "Create target succerfully", "data": data})
}

func (h *Handler) UpdateTarget(c *gin.Context) {
	ok, userJWTProfile := common.ProfileFromJwt(c)
	if !ok {
		c.JSON(common.INTERNAL_SERVER_ERR, gin.H{"error": "User not found"})
		return
	}
	var req models.TargetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(common.INTERNAL_SERVER_ERR, gin.H{"error": err.Error()})
		return
	}
	data, err := h.service.UpdateTarget(c, userJWTProfile.Id, req)
	if err != nil {
		c.JSON(common.INTERNAL_SERVER_ERR, gin.H{"error": err.Error()})
		return
	}
	c.JSON(common.SUCCESS_STATUS, gin.H{"message": "Update target succerfully", "data": data})
}
