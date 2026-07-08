package notification

import (
	"context"
	"kulkasku/internal/helper"
	"net/http"
)

func (r *notificationService) MarkAsRead(ctx context.Context, notificationId, userId int64) (*helper.WebResponse, error) {
	err := r.notificationRepository.MarkAsRead(ctx, notificationId, userId)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   nil,
	}, nil
}
