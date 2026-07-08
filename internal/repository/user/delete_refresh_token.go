package user

import (
	"context"
	"errors"
)

func (r *userRepository) DeleteRefreshToken(ctx context.Context, userId int64) error {
	query := `DELETE FROM refresh_tokens WHERE user_id = ?`

	result, err := r.db.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("Nothing To Delete")
	}

	return nil
}
