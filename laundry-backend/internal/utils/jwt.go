package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("laundry-secret-key")

// Claims represents the structure of our JWT claims
type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT creates a new JWT token with the specified claims
func GenerateJWT(userID, referenceID int, username, role, referenceLevel string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":         userID,         //userAccess.ID,
		"reference_id":    referenceID,    //userAccess.ReferenceID,
		"reference_level": referenceLevel, //userAccess.ReferenceID,
		"username":        username,       //userAccess.Username,
		"role":            role,
		"exp":             float64(time.Now().Add(30 * time.Minute).Unix()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT validates a JWT token and returns the claims
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// ParseJWTClaims parses a JWT token and returns all claims as a map
// This is useful when you want to access custom claims that aren't defined in the Claims struct
func ParseJWTClaims(tokenString string) (jwt.MapClaims, error) {
	token, _, err := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}

// GetClaimValue extracts a specific claim value from a JWT token
// This is a convenience function for accessing individual claims
func GetClaimValue(tokenString string, claimName string) (interface{}, bool) {
	claims, err := ParseJWTClaims(tokenString)
	if err != nil {
		return nil, false
	}

	value, exists := claims[claimName]
	return value, exists
}

// GetStringClaim extracts a string claim value from a JWT token
func GetStringClaim(tokenString string, claimName string) (string, bool) {
	value, exists := GetClaimValue(tokenString, claimName)
	if !exists {
		return "", false
	}

	if str, ok := value.(string); ok {
		return str, true
	}

	return "", false
}

// GetIntClaim extracts an integer claim value from a JWT token
func GetIntClaim(tokenString string, claimName string) (int, bool) {
	value, exists := GetClaimValue(tokenString, claimName)
	if !exists {
		return 0, false
	}

	switch v := value.(type) {
	case float64:
		return int(v), true
	case int:
		return v, true
	case int64:
		return int(v), true
	default:
		return 0, false
	}
}
