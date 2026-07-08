package food

import (
	"context"
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	"kulkasku/internal/model"
	"net/http"
	"time"
)

func (r *foodService) Create(ctx context.Context, createFoodRequest *dto.CreateFoodRequest, userId int64) (*helper.WebResponse, error) {
	shelfLifeDays, err := helper.GetShelfLifeDays(createFoodRequest.Name, createFoodRequest.Category)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error At Get Shelf Life Days",
			Data:   nil,
		}, nil
	}

	now := time.Now()
	expiredAt := now.AddDate(0, 0, shelfLifeDays)

	foodModel := &model.Food{
		UserID:        userId,
		Name:          createFoodRequest.Name,
		Category:      createFoodRequest.Category,
		ImageURL:      createFoodRequest.ImageURL,
		ShelfLifeDays: shelfLifeDays,
		ExpiredAt:     expiredAt,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	foodID, err := r.foodRepository.CreateFood(ctx, foodModel)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error At Create Food",
			Data:   err.Error(),
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Berhasil Create Food",
		Data:   foodID,
	}, nil
}
