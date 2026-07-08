package recipe

import (
	"context"
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	foodRepository "kulkasku/internal/repository/food"
	recipeRepository "kulkasku/internal/repository/recipe"
)

type RecipeService interface {
	Generate(ctx context.Context, userId int64) (*helper.WebResponse, error)
	Save(ctx context.Context, req *dto.SaveRecipeRequest, userId int64) (*helper.WebResponse, error)
}

type recipeService struct {
	foodRepository   foodRepository.FoodRepository
	recipeRepository recipeRepository.RecipeRepository
}

func NewService(foodRepository foodRepository.FoodRepository, recipeRepository recipeRepository.RecipeRepository) RecipeService {
	return &recipeService{
		foodRepository:   foodRepository,
		recipeRepository: recipeRepository,
	}
}
