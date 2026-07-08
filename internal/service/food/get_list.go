package food

import (
	"context"
	"kulkasku/internal/helper"
	"net/http"
)

func (r *foodService) GetList(ctx context.Context, userId int64) (*helper.WebResponse, error) {
	foods, err := r.foodRepository.GetFoodsByUserId(ctx, userId)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error At Get List Food",
			Data:   nil,
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   foods,
	}, nil
}
