package recipe

import (
	"context"
	"kulkasku/internal/model"
)

func (r *recipeRepository) SaveRecipe(ctx context.Context, recipe *model.Recipe) (int64, error) {
	query := `INSERT INTO recipes (user_id, title, description, cooking_time, ingredients_used, missing_ingredients, instructions) VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query, recipe.UserID, recipe.Title, recipe.Description, recipe.CookingTime, recipe.IngredientsUsed, recipe.MissingIngredients, recipe.Instructions)
	if err != nil {
		return 0, err
	}

	recipeID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return recipeID, nil
}
