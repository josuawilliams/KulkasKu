package recipe

import (
	"context"
	"kulkasku/internal/helper"
	"net/http"
)

func (r *recipeService) Delete(ctx context.Context, recipeId, userId int64) (*helper.WebResponse, error) {
	err := r.recipeRepository.DeleteRecipe(ctx, recipeId, userId)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Delete Recipe",
		Data:   nil,
	}, nil
}
