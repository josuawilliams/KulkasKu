package user

import (
	"context"
	"kulkasku/internal/model"
	"time"
)

func (r *userRepository) GetRefreshToken(ctx context.Context, userId int64, now time.Time) (*model.RefreshToken, error) {
	query := `SELECT id, user_id, refresh_token, expires_at FROM refresh_tokens WHERE user_id = ? AND expires_at >= ? `

	row := r.db.QueryRowContext(ctx, query, userId, now)
	refreshToken := &model.RefreshToken{}
	err := row.Scan(&refreshToken.ID, &refreshToken.UserID, &refreshToken.RefreshToken, &refreshToken.ExpiresAt)
	if err != nil {
		return nil, err
	}

	return refreshToken, nil
}
