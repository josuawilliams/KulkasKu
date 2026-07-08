package food

import (
	"context"
	"kulkasku/internal/model"
)

func (r *foodRepository) GetFoodsByUserId(ctx context.Context, userId int64) ([]*model.Food, error) {
	query := `SELECT id, user_id, name, category, image_url, shelf_life_days, expired_at, created_at, updated_at FROM foods WHERE user_id = ?`

	rows, err := r.db.QueryContext(ctx, query, userId)
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
