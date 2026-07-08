package notification

import (
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GenerateRecipe(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req = &dto.CreateNotificationRequest{}
	)

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   nil,
		})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	userId := c.GetInt64("userId")
	res, err := h.notificationService.GenerateRecipe(ctx, req, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   nil,
		})
		return
	}

	c.JSON(res.Code, res)
}
