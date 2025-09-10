package delivery

import (
	"net/http"
	"laundry-backend/internal/entities"
	"laundry-backend/internal/usecases"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authUsecase usecases.AuthUsecase
}

func NewAuthHandler(authUsecase usecases.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var request entities.LoginRequest
	if err := c.Bind(&request); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	response, err := h.authUsecase.Login(request)
	if err != nil {
		return ErrorResponse(c, http.StatusUnauthorized, "Authentication failed", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Login successful", response)
}