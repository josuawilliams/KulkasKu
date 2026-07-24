package user

import (
	"kulkasku/internal/middleware"
	"kulkasku/internal/service/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api         *gin.Engine
	validate    *validator.Validate
	userService user.UserService
}

func NewHandler(api *gin.Engine, validate *validator.Validate, userService user.UserService) Handler {
	return Handler{
		api:         api,
		validate:    validate,
		userService: userService,
	}
}

func (h *Handler) RouteList(secretKey string) {
	authRoute := h.api.Group("/auth")
	authRoute.POST("/register", h.RouteRegister)
	authRoute.POST("/login/google", h.RouteLoginGoogle)
	authRoute.POST("/login", h.Login)

	refreshTokenRoute := h.api.Group("/auth")
	refreshTokenRoute.Use(middleware.AuthRefreshTokenMiddleware(secretKey))
	refreshTokenRoute.POST("/refresh-token", h.RefreshToken)

	authRoute.Use(middleware.AuthMiddleware(secretKey))
	authRoute.PUT("/password", h.UpdatePassword)
	authRoute.GET("/me", h.Me)
}
