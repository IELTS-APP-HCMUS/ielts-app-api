package handlers

import (
	"ielts-app-api/common"
	"ielts-app-api/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetVocab(c *gin.Context) {
	ok, userJWTProfile := common.ProfileFromJwt(c)
	if !ok {
		common.AbortWithError(c, common.ErrUserNotFound)
		return
	}
	data, err := h.service.GetVocabById(c, userJWTProfile.Id)
	if err != nil {
		common.AbortWithError(c, err)
		return
	}
	c.JSON(common.SUCCESS_STATUS, common.BaseResponseMess(common.SUCCESS_STATUS, "Get vocab bank successfully", data))
}

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
	c.JSON(common.SUCCESS_STATUS, common.BaseResponseMess(common.SUCCESS_STATUS, "Create vocab successfully", data))
}

func (h *Handler) UpdateVocab(c *gin.Context) {
	ok, userJWTProfile := common.ProfileFromJwt(c)
	if !ok {
		common.AbortWithError(c, common.ErrUserNotFound)
		return
	}
	var paramsUri models.VocabQuery
	if err := c.ShouldBindQuery(&paramsUri); err != nil {
		common.AbortWithError(c, err)
		return
	}
	var req models.VocabRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.AbortWithError(c, common.ErrInvalidInput)
		return
	}
	data, err := h.service.UpdateVocab(c, userJWTProfile.Id, paramsUri, req)
	if err != nil {
		common.AbortWithError(c, err)
		return
	}
	c.JSON(common.SUCCESS_STATUS, common.BaseResponseMess(common.SUCCESS_STATUS, "Update vocab successfully", data))
}

func (h *Handler) DeleteVocab(c *gin.Context) {
	ok, userJWTProfile := common.ProfileFromJwt(c)
	if !ok {
		common.AbortWithError(c, common.ErrUserNotFound)
		return
	}
	var paramsUri models.VocabQuery
	if err := c.ShouldBindQuery(&paramsUri); err != nil {
		common.AbortWithError(c, err)
		return
	}
	err := h.service.DeleteVocab(c, userJWTProfile.Id, paramsUri)
	if err != nil {
		common.AbortWithError(c, err)
		return
	}
	c.JSON(common.SUCCESS_STATUS, common.BaseResponseMess(common.SUCCESS_STATUS, "Delete vocab successfully", nil))
}

func (h *Handler) GetReadingVocab() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.LookUpVocabRequest
		if err := c.ShouldBindQuery(&request); err != nil {
			common.AbortWithError(c, common.ErrInvalidInput)
			return
		}

		data, err := h.service.GetReadingVocab(c.Request.Context(), request)
		if err != nil {
			common.AbortWithError(c, err)
			return
		}

		c.JSON(common.SUCCESS_STATUS,
			common.BaseResponseMess(common.SUCCESS_STATUS, "Get reading vocab bank successfully", data))
	}
}
