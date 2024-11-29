package handlers

import (
	"ielts-app-api/common"
	"ielts-app-api/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetQuizzes() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID string
		ok, userJWTProfile := common.ProfileFromJwt(c)
		if ok {
			userID = userJWTProfile.Id
		}

		var params = models.ListQuizzesParamsUri{}
		if err := c.ShouldBindQuery(&params); err != nil {
			c.AbortWithStatusJSON(common.BAD_REQUEST_STATUS, common.BaseResponse(common.REQUEST_FAILED, "Thông tin không hợp lệ", err.Error(), nil))
			return
		}

		data, err := h.service.GetQuizzes(c, userID, &params)
		if err != nil {
			c.AbortWithStatusJSON(common.INTERNAL_SERVER_ERR, common.BaseResponse(common.REQUEST_FAILED, "Đã xảy ra lỗi!", err.Error(), nil))
			return
		}

		c.JSON(common.SUCCESS_STATUS, common.ResponseOk(data))
	}
}
