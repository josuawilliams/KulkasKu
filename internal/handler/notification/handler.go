package notification

import (
	"kulkasku/internal/middleware"
	notificationService "kulkasku/internal/service/notification"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api                *gin.Engine
	validate           *validator.Validate
	notificationService notificationService.NotificationService
}

func NewHandler(api *gin.Engine, validate *validator.Validate, notificationService notificationService.NotificationService) Handler {
	return Handler{
		api:                api,
		validate:           validate,
		notificationService: notificationService,
	}
}

func (h *Handler) RouteList(secretKey string) {
	notifRoute := h.api.Group("/notifications")
	notifRoute.Use(middleware.AuthMiddleware(secretKey))
	notifRoute.POST("/generate-expired", h.GenerateExpired)
	notifRoute.POST("/generate-recipe", h.GenerateRecipe)
	notifRoute.GET("/list", h.GetList)
	notifRoute.PUT("/read/:id", h.MarkAsRead)
}
