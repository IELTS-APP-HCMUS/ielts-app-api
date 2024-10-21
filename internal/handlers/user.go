package handlers

import (
	"ielts-app-api/common"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUser(c *gin.Context) {
	// userId := c.Param("user_id")
	userId, exists := c.Get(common.UserId)
	if !exists {
		c.JSON(common.INTERNAL_SERVER_ERR, gin.H{"error": "User not found"})
		return
	}
	// userId = "417b8399-d6a6-49f9-bec1-440ffe5a8f46"
	userIdStr, ok := userId.(string)
	if !ok {
		c.JSON(common.INTERNAL_SERVER_ERR, gin.H{"error": "Invalid user ID type"})
		return
	}
	data, err := h.service.GetUserById(c, userIdStr)
	if err != nil {
		c.JSON(common.INTERNAL_SERVER_ERR, gin.H{"error": err.Error()})
		return
	}
	c.JSON(common.SUCCESS_STATUS, gin.H{"message": "Get user succerfully", "data": data})
}
