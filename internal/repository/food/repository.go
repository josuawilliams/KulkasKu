package food

import (
	"context"
	"database/sql"
	"kulkasku/internal/model"
	"time"
)

type FoodRepository interface {
	CreateFood(ctx context.Context, model *model.Food) (int64, error)
	GetFoodsByUserId(ctx context.Context, userId int64) ([]*model.Food, error)
	DeleteFood(ctx context.Context, foodId, userId int64) error
	GetExpiringFoods(ctx context.Context, from, to time.Time) ([]*model.Food, error)
	GetUserIDsWithMinFoods(ctx context.Context, min int) ([]int64, error)
	GetFirstFoodIdByUserId(ctx context.Context, userId int64) (int64, error)
	GetFoodById(ctx context.Context, foodId, userId int64) (*model.Food, error)
}

type foodRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) FoodRepository {
	return &foodRepository{
		db: db,
	}
}
