package utils

import (
	"fmt"
	"laundry-backend/entities"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func LoggMsg(serviceName, message string, err error) {
	if err != nil {
		fmt.Println("")
		log.Printf("[%s] %s - Error: %v", serviceName, message, err)
	} else {
		fmt.Println("")
		log.Printf("[%s] %s", serviceName, message)
	}
}

func QuerySupport(query string) string {
	count := strings.Count(query, "?")
	for i := 0; i < count; i++ {
		query = strings.Replace(query, "?", "$"+strconv.Itoa(i+1), 1)
	}
	return query
}

// SuccessResponse generates a success response with data
func SuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	response := entities.APIResponse{
		Status:  statusCode,
		Message: message,
		Result:  data,
	}
	return c.JSON(statusCode, response)
}

// ErrorResponse generates an error response
func ErrorResponse(c echo.Context, statusCode int, message string, err string) error {
	response := entities.APIResponse{
		Status:  statusCode,
		Message: message,
		Error:   err,
	}
	return c.JSON(statusCode, response)
}

// MessageResponse generates a response with only a message (no data)
func MessageResponse(c echo.Context, statusCode int, message string) error {
	response := entities.APIResponse{
		Status:  statusCode,
		Message: message,
	}
	return c.JSON(statusCode, response)
}
func GenerateJWT(userID, referenceID int, username, role, referenceLevel string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":         userID,         //userAccess.ID,
		"reference_id":    referenceID,    //userAccess.ReferenceID,
		"reference_level": referenceLevel, //userAccess.ReferenceID,
		"username":        username,       //userAccess.Username,
		"role":            role,
		"exp":             float64(time.Now().Add(90 * time.Minute).Unix()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("laundry-secret-key"))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}

// generateInvoiceNumber generates a unique invoice number
func GenerateInvoiceNumber() string {
	// Use current timestamp and random number to generate unique invoice number
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	// Generate a random 4-digit number
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(9000) + 1000

	return fmt.Sprintf("INV%d%02d%02d%02d%02d%02d%d", year, month, day, hour, minute, second, random)
}
