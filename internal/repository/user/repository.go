package user

import (
	"context"
	"database/sql"
	"kulkasku/internal/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, model *model.User) (int64, error)
	GetRefreshToken(ctx context.Context, userId int64, now time.Time) (*model.RefreshToken, error)
	StoreRefreshToken(ctx context.Context, model *model.RefreshToken) error
	GetUserById(ctx context.Context, userId int64) (*model.User, error)
	DeleteRefreshToken(ctx context.Context, userId int64) error
	UpdateUserGoogle(ctx context.Context, userId int64, googleId, avatarURL string) error
	UpdateUserPassword(ctx context.Context, userId int64, passwordHash string) error
}

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
