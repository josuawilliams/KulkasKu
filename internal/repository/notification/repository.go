package notification

import (
	"context"
	"database/sql"
	"kulkasku/internal/model"
)

type NotificationRepository interface {
	Save(ctx context.Context, notification *model.Notification) (int64, error)
	GetByUserId(ctx context.Context, userId int64) ([]*model.Notification, error)
	MarkAsRead(ctx context.Context, notificationId, userId int64) error
}

type notificationRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) NotificationRepository {
	return &notificationRepository{
		db: db,
	}
}
