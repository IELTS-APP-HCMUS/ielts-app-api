package handlers

import (
	"ielts-app-api/common"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMasterData(c *gin.Context) {
	// data, err := h.service.GetMasterData(c)
	// data, err = nil, nil
	// if err != nil {
	// 	common.AbortWithError(c, err)
	// 	return
	// }
	c.JSON(common.SUCCESS_STATUS, common.ResponseOk(123))
}
