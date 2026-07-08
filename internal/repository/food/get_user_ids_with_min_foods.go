package food

import (
	"context"
)

func (r *foodRepository) GetUserIDsWithMinFoods(ctx context.Context, min int) ([]int64, error) {
	query := `SELECT user_id FROM foods GROUP BY user_id HAVING COUNT(*) >= ?`

	rows, err := r.db.QueryContext(ctx, query, min)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userIDs := make([]int64, 0)
	for rows.Next() {
		var userID int64
		if err := rows.Scan(&userID); err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}

	return userIDs, nil
}
