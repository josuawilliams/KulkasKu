package dto

type (
	CreateNotificationRequest struct {
		FoodID  int64  `json:"food_id" validate:"required"`
		Title   string `json:"title" validate:"required"`
		Message string `json:"message" validate:"required"`
		Type    string `json:"type"`
	}

	NotificationResponse struct {
		ID        int64  `json:"id"`
		UserID    int64  `json:"user_id"`
		FoodID    int64  `json:"food_id"`
		Title     string `json:"title"`
		Message   string `json:"message"`
		Type      string `json:"type"`
		IsRead    bool   `json:"is_read"`
		NotifyAt string `json:"notify_at"`
		CreatedAt string `json:"created_at"`
	}
)
