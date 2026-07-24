package user

import (
	"context"
)

func (r *userRepository) UpdateUserPassword(ctx context.Context, userId int64, passwordHash string) error {
	query := `UPDATE users SET password = ? WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, passwordHash, userId)
	return err
}
