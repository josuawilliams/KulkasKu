package recipe

import (
	"kulkasku/internal/middleware"
	recipeService "kulkasku/internal/service/recipe"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api           *gin.Engine
	validate      *validator.Validate
	recipeService recipeService.RecipeService
}

func NewHandler(api *gin.Engine, validate *validator.Validate, recipeService recipeService.RecipeService) Handler {
	return Handler{
		api:           api,
		validate:      validate,
		recipeService: recipeService,
	}
}

func (h *Handler) RouteList(secretKey string) {
	recipeRoute := h.api.Group("/recipes")
	recipeRoute.Use(middleware.AuthMiddleware(secretKey))
	recipeRoute.POST("/generate", h.Generate)
	recipeRoute.POST("/save", h.Save)
	recipeRoute.DELETE("/delete/:id", h.Delete)
}
