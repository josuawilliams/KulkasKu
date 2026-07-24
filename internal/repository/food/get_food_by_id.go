package food

import (
	"context"
	"kulkasku/internal/model"
)

func (r *foodRepository) GetFoodById(ctx context.Context, foodId, userId int64) (*model.Food, error) {
	query := `SELECT id, user_id, name, category, image_url, shelf_life_days, expired_at, created_at, updated_at FROM foods WHERE id = ? AND user_id = ?`

	food := &model.Food{}
	err := r.db.QueryRowContext(ctx, query, foodId, userId).Scan(&food.ID, &food.UserID, &food.Name, &food.Category, &food.ImageURL, &food.ShelfLifeDays, &food.ExpiredAt, &food.CreatedAt, &food.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return food, nil
}
