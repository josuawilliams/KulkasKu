package model

import "time"

type Recipe struct {
	ID                 int64     `json:"id"`
	UserID             int64     `json:"user_id"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	CookingTime        int       `json:"cooking_time"`
	IngredientsUsed    string    `json:"ingredients_used"`
	MissingIngredients string    `json:"missing_ingredients"`
	Instructions       string    `json:"instructions"`
	CreatedAt          time.Time `json:"created_at"`
}
