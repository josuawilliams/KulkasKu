package food

import (
	"kulkasku/internal/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetDetail(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetInt64("userId")

	foodId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Invalid Food ID",
			Data:   nil,
		})
		return
	}

	res, err := h.foodService.GetDetail(ctx, foodId, userId)
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
