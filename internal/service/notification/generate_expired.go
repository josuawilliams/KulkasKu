package notification

import (
	"context"
	"kulkasku/internal/helper"
	"kulkasku/internal/model"
	"net/http"
	"time"
)

func (r *notificationService) GenerateExpired(ctx context.Context, userId int64) (*helper.WebResponse, error) {
	foods, err := r.foodRepository.GetFoodsByUserId(ctx, userId)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Gagal mengambil data makanan",
			Data:   nil,
		}, nil
	}

	now := time.Now()
	threeDaysLater := now.AddDate(0, 0, 3)
	created := make([]int64, 0)

	for _, food := range foods {
		if food.ExpiredAt.After(now) && food.ExpiredAt.Before(threeDaysLater) {
			notifID, err := r.notificationRepository.Save(ctx, &model.Notification{
				UserID:   userId,
				FoodID:   food.ID,
				Title:    "Makanan Akan Kadaluarsa",
				Message:  food.Name + " akan kadaluarsa pada " + food.ExpiredAt.Format("02 Jan 2006"),
				Type:     "expired",
				IsRead:   false,
				NotifyAt: food.ExpiredAt,
			})
			if err != nil {
				return &helper.WebResponse{
					Code:   http.StatusInternalServerError,
					Status: "Gagal menyimpan notifikasi",
					Data:   nil,
				}, nil
			}
			created = append(created, notifID)
		}
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   created,
	}, nil
}
