package user

import (
	"context"
	"kulkasku/internal/model"
)

func (r *userRepository) CreateUser(ctx context.Context, model *model.User) (int64, error) {
	query := `INSERT INTO users (name, email, password, google_id, avatar_url) VALUES (?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query, model.Name, model.Email, model.Password, model.GoogleID, model.AvatarURL)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil

}
