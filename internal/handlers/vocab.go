package handlers

import (
	"ielts-app-api/common"
	"ielts-app-api/internal/models"

	"github.com/gin-gonic/gin"
)

// func (h *Handler) GetTarget(c *gin.Context) {
// 	ok, userJWTProfile := common.ProfileFromJwt(c)
// 	if !ok {
// 		common.AbortWithError(c, common.ErrUserNotFound)
// 		return
// 	}
// 	data, err := h.service.GetTargetById(c, userJWTProfile.Id)
// 	if err != nil {
// 		common.AbortWithError(c, err)
// 		return
// 	}
// 	c.JSON(common.SUCCESS_STATUS, common.BaseResponseMess(common.SUCCESS_STATUS, "Get target successfully", data))
// }

func (h *Handler) CreateVocab(c *gin.Context) {
	ok, userJWTProfile := common.ProfileFromJwt(c)
	if !ok {
		common.AbortWithError(c, common.ErrUserNotFound)
		return
	}
	var req models.VocabRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.AbortWithError(c, common.ErrInvalidInput)
		return
	}
	data, err := h.service.CreateVocab(c, userJWTProfile.Id, req)
	if err != nil {
		common.AbortWithError(c, err)
		return
	}
	c.JSON(common.SUCCESS_STATUS, common.BaseResponseMess(common.SUCCESS_STATUS, "Create target successfully", data))
}