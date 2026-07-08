package notification

import (
	"kulkasku/internal/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) MarkAsRead(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetInt64("userId")

	notifID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Invalid Notification ID",
			Data:   nil,
		})
		return
	}

	res, err := h.notificationService.MarkAsRead(ctx, notifID, userId)
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
