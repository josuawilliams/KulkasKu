package food

import (
	"context"
)

func (r *foodRepository) GetFirstFoodIdByUserId(ctx context.Context, userId int64) (int64, error) {
	query := `SELECT id FROM foods WHERE user_id = ? LIMIT 1`

	var foodID int64
	err := r.db.QueryRowContext(ctx, query, userId).Scan(&foodID)
	if err != nil {
		return 0, err
	}

	return foodID, nil
}
