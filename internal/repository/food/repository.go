package food

import (
	"context"
	"database/sql"
	"kulkasku/internal/model"
)

type FoodRepository interface {
	CreateFood(ctx context.Context, model *model.Food) (int64, error)
	GetFoodsByUserId(ctx context.Context, userId int64) ([]*model.Food, error)
	DeleteFood(ctx context.Context, foodId, userId int64) error
}

type foodRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) FoodRepository {
	return &foodRepository{
		db: db,
	}
}
