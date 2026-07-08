package user

import (
	"context"
	"database/sql"
	"kulkasku/internal/model"
)

func (r *userRepository) GetUserById(ctx context.Context, userId int64) (*model.User, error) {
	query := `SELECT id, name, email, password, google_id, avatar_url, created_at, updated_at FROM users WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, userId)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.GoogleID, &user.AvatarURL, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
