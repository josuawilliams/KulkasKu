package notification

import (
	"context"
	"database/sql"
	"kulkasku/internal/model"
)

func (r *notificationRepository) GetByUserId(ctx context.Context, userId int64) ([]*model.Notification, error) {
	query := `SELECT id, user_id, food_id, title, message, type, is_read, notify_at, created_at, updated_at FROM notifications WHERE user_id = ? ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

		notifications := make([]*model.Notification, 0)
	for rows.Next() {
		n := &model.Notification{}
		var message, nType sql.NullString
		var notifyAt, updatedAt sql.NullTime

		err := rows.Scan(&n.ID, &n.UserID, &n.FoodID, &n.Title, &message, &nType, &n.IsRead, &notifyAt, &n.CreatedAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		n.Message = message.String
		n.Type = nType.String
		n.NotifyAt = notifyAt.Time
		n.UpdatedAt = updatedAt.Time

		notifications = append(notifications, n)
	}

	return notifications, nil
}
