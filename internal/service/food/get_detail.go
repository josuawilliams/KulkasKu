package food

import (
	"context"
	"database/sql"
	"errors"
	"kulkasku/internal/helper"
	"net/http"
)

func (r *foodService) GetDetail(ctx context.Context, foodId, userId int64) (*helper.WebResponse, error) {
	food, err := r.foodRepository.GetFoodById(ctx, foodId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &helper.WebResponse{
				Code:   http.StatusNotFound,
				Status: "Food Not Found",
				Data:   nil,
			}, nil
		}
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   nil,
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   food,
	}, nil
}
