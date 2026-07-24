package user

import (
	"context"
	"database/sql"
)

func (r *userRepository) UpdateUserGoogle(ctx context.Context, userId int64, googleId, avatarURL string) error {
	query := `UPDATE users SET google_id = ?, avatar_url = ? WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query, googleId, avatarURL, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
