package recipe

import (
	"context"
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	"kulkasku/internal/model"
	"net/http"
	"strings"
	"time"
)

func (r *recipeService) Save(ctx context.Context, req *dto.SaveRecipeRequest, userId int64) (*helper.WebResponse, error) {
	now := time.Now()
	recipe := &model.Recipe{
		UserID:             userId,
		Title:              req.Title,
		Description:        req.Description,
		CookingTime:        req.CookingTime,
		IngredientsUsed:    req.IngredientsUsed,
		MissingIngredients: req.MissingIngredients,
		Instructions:       strings.TrimSpace(req.Instructions),
		CreatedAt:          now,
	}

	recipeID, err := r.recipeRepository.SaveRecipe(ctx, recipe)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Gagal menyimpan resep",
			Data:   nil,
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Save Recipe",
		Data:   recipeID,
	}, nil
}
