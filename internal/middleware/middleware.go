package middleware

import (
	"kulkasku/internal/helper"
	"kulkasku/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		if header == "" {
			c.JSON(http.StatusUnauthorized, helper.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized (19)",
				Data:   nil,
			})
			c.Abort()
			return
		}

		userId, email, err := jwt.ValidateToken(header, secretKey, true)
		if err != nil {
			c.JSON(http.StatusUnauthorized, helper.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized (31)",
				Data:   nil,
			})
			c.Abort()
			return
		}

		c.Set("userId", userId)
		c.Set("email", email)
		c.Next()

	}
}

func AuthRefreshTokenMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		if header == "" {
			c.JSON(http.StatusUnauthorized, helper.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized (53)",
				Data:   nil,
			})
			c.Abort()
			return
		}

		userId, email, err := jwt.ValidateToken(header, secretKey, false)
		if err != nil {
			c.JSON(http.StatusUnauthorized, helper.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized (64)",
				Data:   nil,
			})
			c.Abort()
			return
		}

		c.Set("userId", userId)
		c.Set("email", email)
		c.Next()

	}
}
