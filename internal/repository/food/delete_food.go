package food

import (
	"context"
	"errors"
)

func (r *foodRepository) DeleteFood(ctx context.Context, foodId, userId int64) error {
	query := `DELETE FROM foods WHERE id = ? AND user_id = ?`

	result, err := r.db.ExecContext(ctx, query, foodId, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("food not found")
	}

	return nil
}
