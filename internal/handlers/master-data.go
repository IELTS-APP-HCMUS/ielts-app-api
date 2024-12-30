package handlers

import (
	"ielts-app-api/common"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMasterData(c *gin.Context) {
	c.JSON(common.SUCCESS_STATUS, common.ResponseOk(123))
}
