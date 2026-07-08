package user

import (
	"context"
	"kulkasku/internal/model"
)

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, name, email, password, google_id, avatar_url, created_at, updated_at FROM users WHERE email = ? LIMIT 1`
	
	row := r.db.QueryRowContext(ctx, query, email)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.GoogleID, &user.AvatarURL, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}
