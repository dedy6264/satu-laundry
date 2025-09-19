package delivery

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/usecases"
	"laundry-backend/internal/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ServiceCategoryHandler struct {
	serviceCategoryUsecase usecases.ServiceCategoryUsecase
}

func NewServiceCategoryHandler(serviceCategoryUsecase usecases.ServiceCategoryUsecase) *ServiceCategoryHandler {
	return &ServiceCategoryHandler{
		serviceCategoryUsecase: serviceCategoryUsecase,
	}
}

func (h *ServiceCategoryHandler) CreateServiceCategory(c echo.Context) error {
	var (
		request entities.CreateServiceCategoryRequest
		svcName = "CreateServiceCategory"
	)

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	request.Name = strings.ToUpper(request.Name)

	if err := h.serviceCategoryUsecase.CreateServiceCategory(request); err != nil {
		utils.LoggMsg(svcName, "Failed to create service category", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create service category", err.Error())
	}

	return MessageResponse(c, http.StatusCreated, "Service category created successfully")
}

func (h *ServiceCategoryHandler) GetServiceCategoryByID(c echo.Context) error {
	var (
		svcName = "GetServiceCategoryByID"
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid service category ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid service category ID", err.Error())
	}

	category, err := h.serviceCategoryUsecase.GetServiceCategoryByID(id)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get service category", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get service category", err.Error())
	}

	if category == nil {
		utils.LoggMsg(svcName, "Service category not found", nil)
		return ErrorResponse(c, http.StatusNotFound, "Service category not found", "Service category with given ID does not exist")
	}

	return SuccessResponse(c, http.StatusOK, "Service category retrieved successfully", category)
}

func (h *ServiceCategoryHandler) GetAllServiceCategories(c echo.Context) error {
	var (
		svcName = "GetAllServiceCategories"
		request entities.DataTablesRequest
	)

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	response, err := h.serviceCategoryUsecase.GetAllServiceCategoriesDataTables(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get service categories", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get service categories", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Service categories retrieved successfully", response)
}

func (h *ServiceCategoryHandler) UpdateServiceCategory(c echo.Context) error {
	var (
		svcName = "UpdateServiceCategory"
		request entities.UpdateServiceCategoryRequest
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid service category ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid service category ID", err.Error())
	}

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	request.Name = strings.ToUpper(request.Name)

	if err := h.serviceCategoryUsecase.UpdateServiceCategory(id, request); err != nil {
		utils.LoggMsg(svcName, "Failed to update service category", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update service category", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Service category updated successfully")
}

func (h *ServiceCategoryHandler) DeleteServiceCategory(c echo.Context) error {
	var (
		svcName = "DeleteServiceCategory"
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid service category ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid service category ID", err.Error())
	}

	if err := h.serviceCategoryUsecase.DeleteServiceCategory(id); err != nil {
		utils.LoggMsg(svcName, "Failed to delete service category", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete service category", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Service category deleted successfully")
}