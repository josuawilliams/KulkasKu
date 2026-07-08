package notification

import (
	"context"
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	"kulkasku/internal/model"
	"net/http"
	"time"
)

func (r *notificationService) GenerateRecipe(ctx context.Context, req *dto.CreateNotificationRequest, userId int64) (*helper.WebResponse, error) {
	if req.Type == "" {
		req.Type = "recipe"
	}

	notifID, err := r.notificationRepository.Save(ctx, &model.Notification{
		UserID:   userId,
		FoodID:   req.FoodID,
		Title:    req.Title,
		Message:  req.Message,
		Type:     req.Type,
		IsRead:   false,
		NotifyAt: time.Now(),
	})
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Gagal menyimpan notifikasi",
			Data:   nil,
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   notifID,
	}, nil
}
