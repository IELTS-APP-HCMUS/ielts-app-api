package handlers

import (
	"ielts-app-api/common"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserProfile(c *gin.Context) {
	ok, userJWTProfile := common.ProfileFromJwt(c)
	if !ok {
		common.AbortWithError(c, common.ErrUserNotFound)
		return
	}
	data, err := h.service.GetUserProfileById(c, userJWTProfile.Id)
	if err != nil {
		common.AbortWithError(c, err)
		return
	}
	c.JSON(common.SUCCESS_STATUS, common.BaseResponseMess(common.SUCCESS_STATUS, "Get user profile successfully", data))
}
