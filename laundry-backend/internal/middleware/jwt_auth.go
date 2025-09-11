package middleware

import (
	"fmt"
	"laundry-backend/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// JWTAuth is a middleware that validates JWT tokens and extracts claims
func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("=== JWT MIDDLEWARE START ===\n")
		// Get the Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		fmt.Printf("Authorization header: %s\n", authHeader)
		if authHeader == "" {
			fmt.Printf("Missing authorization header\n")
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Missing authorization header",
			})
		}

		// Check if the header has the correct format "Bearer <token>"
		const bearerPrefix = "Bearer "
		if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
			fmt.Printf("Invalid authorization header format\n")
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid authorization header format",
			})
		}

		// Extract the token
		tokenString := authHeader[len(bearerPrefix):]
		fmt.Printf("Token extracted: %s\n", tokenString[:min(20, len(tokenString))] + "...")

		// First, try to validate with our standard claims structure
		claims, err := utils.ValidateJWT(tokenString)
		if err == nil {
			fmt.Printf("Standard validation successful\n")
			// Successfully validated with standard claims
			c.Set("user_id", claims.UserID)
			c.Set("email", claims.Email)
			c.Set("role", claims.Role)
			fmt.Printf("Set context values - user_id: %d, email: %s, role: %s\n", claims.UserID, claims.Email, claims.Role)
			fmt.Printf("=== JWT MIDDLEWARE END (SUCCESS) ===\n")
			return next(c)
		}
		fmt.Printf("Standard validation failed: %v\n", err)

		// If standard validation fails, try to parse as generic claims
		_, err = utils.ParseJWTClaims(tokenString)
		if err != nil {
			fmt.Printf("Generic parsing failed: %v\n", err)
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid or expired token",
			})
		}
		fmt.Printf("Generic parsing successful\n")

		// Extract common claims from the token string directly
		if userID, ok := utils.GetIntClaim(tokenString, "user_id"); ok {
			c.Set("user_id", userID)
			fmt.Printf("Set user_id from generic claims: %d\n", userID)
		} else {
			fmt.Printf("Failed to extract user_id from generic claims\n")
		}

		if email, ok := utils.GetStringClaim(tokenString, "email"); ok {
			c.Set("email", email)
			fmt.Printf("Set email from generic claims: %s\n", email)
		}

		if role, ok := utils.GetStringClaim(tokenString, "role"); ok {
			c.Set("role", role)
			fmt.Printf("Set role from generic claims: %s\n", role)
		}
		
		fmt.Printf("=== JWT MIDDLEWARE END (GENERIC SUCCESS) ===\n")
		return next(c)
	}
}

// Helper function to get minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
