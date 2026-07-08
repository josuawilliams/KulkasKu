package dto

type CreateFoodRequest struct {
	Name     string `json:"name" validate:"required"`
	Category string `json:"category" validate:"required"`
	ImageURL string `json:"image_url" validate:"required"`
}
