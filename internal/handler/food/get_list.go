package food

import (
	"kulkasku/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetList(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetInt64("userId")

	res, err := h.foodService.GetList(ctx, userId)
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
