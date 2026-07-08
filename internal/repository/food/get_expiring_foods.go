package food

import (
	"context"
	"kulkasku/internal/model"
	"time"
)

func (r *foodRepository) GetExpiringFoods(ctx context.Context, from, to time.Time) ([]*model.Food, error) {
	query := `SELECT id, user_id, name, category, image_url, shelf_life_days, expired_at, created_at, updated_at FROM foods WHERE expired_at >= ? AND expired_at <= ?`

	rows, err := r.db.QueryContext(ctx, query, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	foods := make([]*model.Food, 0)
	for rows.Next() {
		food := &model.Food{}
		err := rows.Scan(&food.ID, &food.UserID, &food.Name, &food.Category, &food.ImageURL, &food.ShelfLifeDays, &food.ExpiredAt, &food.CreatedAt, &food.UpdatedAt)
		if err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}

	return foods, nil
}
