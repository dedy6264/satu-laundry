package delivery

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/usecases"
	"laundry-backend/internal/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type ServiceHandler struct {
	serviceUsecase usecases.ServiceUsecase
}

func NewServiceHandler(serviceUsecase usecases.ServiceUsecase) *ServiceHandler {
	return &ServiceHandler{
		serviceUsecase: serviceUsecase,
	}
}

func (h *ServiceHandler) CreateService(c echo.Context) error {
	var (
		request entities.CreateServiceRequest
		svcName = "CreateService"
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	request.Name = strings.ToUpper(request.Name)

	if err := h.serviceUsecase.CreateService(request); err != nil {
		utils.LoggMsg(svcName, "Failed to create service", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create service", err.Error())
	}

	return MessageResponse(c, http.StatusCreated, "Service created successfully")
}

func (h *ServiceHandler) GetAllServices(c echo.Context) error {
	var (
		svcName = "GetAllServices"
		request entities.DTRequest[entities.Service]
	)
	// Ambil token dari context
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	request.UserID = int(claims["user_id"].(float64)) // JSON number → float64 → int)
	response, err := h.serviceUsecase.GetAllServicesDataTables(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get services", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get services", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Services retrieved successfully", response)
}

func (h *ServiceHandler) UpdateService(c echo.Context) error {
	var (
		svcName = "UpdateService"
		request entities.Service
	)

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	request.Name = strings.ToUpper(request.Name)

	if err := h.serviceUsecase.UpdateService(request); err != nil {
		utils.LoggMsg(svcName, "Failed to update service", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update service", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Service updated successfully")
}

func (h *ServiceHandler) DeleteService(c echo.Context) error {
	var (
		svcName = "DeleteService"
		request entities.Service
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	if err := h.serviceUsecase.DeleteService(request.ID); err != nil {
		utils.LoggMsg(svcName, "Failed to delete service", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete service", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Service deleted successfully")
}
