package middleware

import (
	"net/http"
	"laundry-backend/internal/utils"

	"github.com/labstack/echo/v4"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Missing authorization header",
			})
		}

		// Format: "Bearer <token>"
		tokenString := authHeader[len("Bearer "):]

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid or expired token",
			})
		}

		// Simpan claims di context
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		return next(c)
	}
}