package recipe

import (
	"context"
	"database/sql"
	"kulkasku/internal/model"
)

type RecipeRepository interface {
	SaveRecipe(ctx context.Context, recipe *model.Recipe) (int64, error)
	DeleteRecipe(ctx context.Context, recipeId, userId int64) error
}

type recipeRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) RecipeRepository {
	return &recipeRepository{
		db: db,
	}
}
