package model

import "time"

type Food struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id"`
	Name          string    `json:"name"`
	Category      string    `json:"category"`
	ImageURL      string    `json:"image_url"`
	ShelfLifeDays int       `json:"shelf_life_days"`
	ExpiredAt     time.Time `json:"expired_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
