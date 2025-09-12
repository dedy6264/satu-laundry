package delivery

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/usecases"
	"laundry-backend/internal/utils"
	"net/http"
	"strconv"

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

	if err := h.serviceUsecase.CreateService(request); err != nil {
		utils.LoggMsg(svcName, "Failed to create service", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create service", err.Error())
	}

	return MessageResponse(c, http.StatusCreated, "Service created successfully")
}

func (h *ServiceHandler) GetServiceByID(c echo.Context) error {
	var (
		svcName = "GetServiceByID"
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid service ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid service ID", err.Error())
	}

	service, err := h.serviceUsecase.GetServiceByID(id)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get service", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get service", err.Error())
	}

	if service == nil {
		utils.LoggMsg(svcName, "Service not found", nil)
		return ErrorResponse(c, http.StatusNotFound, "Service not found", "Service with given ID does not exist")
	}

	return SuccessResponse(c, http.StatusOK, "Service retrieved successfully", service)
}

func (h *ServiceHandler) GetAllServices(c echo.Context) error {
	var (
		svcName = "GetAllServices"
		request entities.DataTablesRequest
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

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
		request entities.UpdateServiceRequest
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid service ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid service ID", err.Error())
	}

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := h.serviceUsecase.UpdateService(id, request); err != nil {
		utils.LoggMsg(svcName, "Failed to update service", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update service", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Service updated successfully")
}

func (h *ServiceHandler) DeleteService(c echo.Context) error {
	var (
		svcName = "DeleteService"
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid service ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid service ID", err.Error())
	}

	if err := h.serviceUsecase.DeleteService(id); err != nil {
		utils.LoggMsg(svcName, "Failed to delete service", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete service", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Service deleted successfully")
}

func (h *ServiceHandler) GetServicesByCategoryID(c echo.Context) error {
	var (
		svcName = "GetServicesByCategoryID"
	)
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid category ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid category ID", err.Error())
	}

	services, err := h.serviceUsecase.GetServicesByCategoryID(categoryID)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get services by category", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get services by category", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Services retrieved successfully", services)
}
