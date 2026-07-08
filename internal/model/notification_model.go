package model

import "time"

type Notification struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	FoodID    int64     `json:"food_id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	IsRead    bool      `json:"is_read"`
	NotifyAt  time.Time `json:"notify_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
