package dto

type (
	RecipeGenerateResponse struct {
		ID                 int64    `json:"id"`
		UserID             int64    `json:"user_id"`
		Title              string   `json:"title"`
		Description        string   `json:"description"`
		CookingTime        int      `json:"cooking_time"`
		IngredientsUsed    string   `json:"ingredients_used"`
		MissingIngredients string   `json:"missing_ingredients"`
		Instructions       []string `json:"instructions"`
		CreatedAt          string   `json:"created_at"`
	}

	SaveRecipeRequest struct {
		Title              string `json:"title" validate:"required"`
		Description        string `json:"description"`
		CookingTime        int    `json:"cooking_time"`
		IngredientsUsed    string `json:"ingredients_used" validate:"required"`
		MissingIngredients string `json:"missing_ingredients"`
		Instructions       string `json:"instructions" validate:"required"`
	}
)
