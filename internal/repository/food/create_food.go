package food

import (
	"context"
	"kulkasku/internal/model"
)

func (r *foodRepository) CreateFood(ctx context.Context, model *model.Food) (int64, error) {
	query := `INSERT INTO foods (user_id, name, category, image_url, shelf_life_days, expired_at) VALUES (?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query, model.UserID, model.Name, model.Category, model.ImageURL, model.ShelfLifeDays, model.ExpiredAt)
	if err != nil {
		return 0, err
	}

	foodID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return foodID, nil
}
