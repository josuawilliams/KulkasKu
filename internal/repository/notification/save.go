package notification

import (
	"context"
	"kulkasku/internal/model"
)

func (r *notificationRepository) Save(ctx context.Context, notification *model.Notification) (int64, error) {
	query := `INSERT INTO notifications (user_id, food_id, title, message, type, is_read, notify_at) VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query, notification.UserID, notification.FoodID, notification.Title, notification.Message, notification.Type, notification.IsRead, notification.NotifyAt)
	if err != nil {
		return 0, err
	}

	notifID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return notifID, nil
}
