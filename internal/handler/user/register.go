package user

import (
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RouteRegister(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req = &dto.RegisterRequest{}
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

	res, err := h.userService.Register(ctx, req)
	if(err != nil){
		c.JSON(http.StatusInternalServerError, helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}
	c.JSON(res.Code, res)

}