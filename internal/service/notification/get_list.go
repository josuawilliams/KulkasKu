package notification

import (
	"context"
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	"net/http"
	"time"
)

func (r *notificationService) GetList(ctx context.Context, userId int64) (*helper.WebResponse, error) {
	notifications, err := r.notificationRepository.GetByUserId(ctx, userId)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Gagal mengambil notifikasi",
			Data:   nil,
		}, nil
	}

	resp := make([]*dto.NotificationResponse, 0, len(notifications))
	for _, n := range notifications {
		resp = append(resp, &dto.NotificationResponse{
			ID:        n.ID,
			UserID:    n.UserID,
			FoodID:    n.FoodID,
			Title:     n.Title,
			Message:   n.Message,
			Type:      n.Type,
			IsRead:    n.IsRead,
			NotifyAt: n.NotifyAt.Format(time.RFC3339),
			CreatedAt: n.CreatedAt.Format(time.RFC3339),
		})
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   resp,
	}, nil
}
