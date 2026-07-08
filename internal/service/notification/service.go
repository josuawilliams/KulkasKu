package notification

import (
	"context"
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	foodRepository "kulkasku/internal/repository/food"
	notificationRepository "kulkasku/internal/repository/notification"
)

type NotificationService interface {
	GenerateExpired(ctx context.Context, userId int64) (*helper.WebResponse, error)
	GenerateRecipe(ctx context.Context, req *dto.CreateNotificationRequest, userId int64) (*helper.WebResponse, error)
	GetList(ctx context.Context, userId int64) (*helper.WebResponse, error)
	MarkAsRead(ctx context.Context, notificationId, userId int64) (*helper.WebResponse, error)
}

type notificationService struct {
	foodRepository         foodRepository.FoodRepository
	notificationRepository notificationRepository.NotificationRepository
}

func NewService(foodRepository foodRepository.FoodRepository, notificationRepository notificationRepository.NotificationRepository) NotificationService {
	return &notificationService{
		foodRepository:         foodRepository,
		notificationRepository: notificationRepository,
	}
}
