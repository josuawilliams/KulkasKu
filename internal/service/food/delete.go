package food

import (
	"context"
	"kulkasku/internal/helper"
	"net/http"
)

func (r *foodService) Delete(ctx context.Context, foodId, userId int64) (*helper.WebResponse, error) {
	err := r.foodRepository.DeleteFood(ctx, foodId, userId)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Delete Food",
		Data:   nil,
	}, nil
}
