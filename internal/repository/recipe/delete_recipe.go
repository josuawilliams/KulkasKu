package recipe

import (
	"context"
	"errors"
)

func (r *recipeRepository) DeleteRecipe(ctx context.Context, recipeId, userId int64) error {
	query := `DELETE FROM recipes WHERE id = ? AND user_id = ?`

	result, err := r.db.ExecContext(ctx, query, recipeId, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("recipe not found")
	}

	return nil
}
