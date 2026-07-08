package recipe

import (
	"context"
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	"net/http"
	"time"
)

func (r *recipeService) Generate(ctx context.Context, userId int64) (*helper.WebResponse, error) {
	foods, err := r.foodRepository.GetFoodsByUserId(ctx, userId)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error At Get Foods",
			Data:   nil,
		}, nil
	}

	if len(foods) < 5 {
		return &helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Minimal 5 bahan makanan diperlukan untuk generate resep",
			Data:   nil,
		}, nil
	}

	foodNames := make([]string, 0, len(foods))
	for _, f := range foods {
		foodNames = append(foodNames, f.Name)
	}

	aiRecipes, err := helper.GenerateRecipesFromAI(foodNames)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Gagal generate resep dari AI",
			Data:   nil,
		}, nil
	}

	now := time.Now().Format(time.RFC3339)
	recipes := make([]*dto.RecipeGenerateResponse, 0, len(aiRecipes))
	for i, r := range aiRecipes {
		recipes = append(recipes, &dto.RecipeGenerateResponse{
			ID:                 int64(i + 1),
			UserID:             userId,
			Title:              r.Title,
			Description:        r.Description,
			CookingTime:        r.CookingTime,
			IngredientsUsed:    r.IngredientsUsed,
			MissingIngredients: r.MissingIngredients,
			Instructions:       r.Instructions,
			CreatedAt:          now,
		})
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   recipes,
	}, nil
}
