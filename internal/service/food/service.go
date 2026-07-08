package food

import (
	"context"
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	foodRepository "kulkasku/internal/repository/food"
)

type FoodService interface {
	Create(ctx context.Context, createFoodRequest *dto.CreateFoodRequest, userId int64) (*helper.WebResponse, error)
	GetList(ctx context.Context, userId int64) (*helper.WebResponse, error)
	Delete(ctx context.Context, foodId, userId int64) (*helper.WebResponse, error)
}

type foodService struct {
	foodRepository foodRepository.FoodRepository
}

func NewService(foodRepository foodRepository.FoodRepository) FoodService {
	return &foodService{
		foodRepository: foodRepository,
	}
}
