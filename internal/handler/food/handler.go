package food

import (
	"kulkasku/internal/middleware"
	foodService "kulkasku/internal/service/food"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api         *gin.Engine
	validate    *validator.Validate
	foodService foodService.FoodService
}

func NewHandler(api *gin.Engine, validate *validator.Validate, foodService foodService.FoodService) Handler {
	return Handler{
		api:         api,
		validate:    validate,
		foodService: foodService,
	}
}

func (h *Handler) RouteList(secretKey string) {
	foodRoute := h.api.Group("/foods")
	foodRoute.Use(middleware.AuthMiddleware(secretKey))
	foodRoute.POST("/create", h.CreateFood)
	foodRoute.GET("/list", h.GetList)
	foodRoute.GET("/detail/:id", h.GetDetail)
	foodRoute.DELETE("/delete/:id", h.Delete)
}
