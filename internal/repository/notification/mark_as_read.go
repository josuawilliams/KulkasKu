package notification

import (
	"context"
	"errors"
)

func (r *notificationRepository) MarkAsRead(ctx context.Context, notificationId, userId int64) error {
	query := `UPDATE notifications SET is_read = TRUE WHERE id = ? AND user_id = ?`

	result, err := r.db.ExecContext(ctx, query, notificationId, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("notification not found")
	}

	return nil
}
