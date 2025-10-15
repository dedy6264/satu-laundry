package delivery

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/usecases"
	"laundry-backend/internal/utils"
	"net/http"
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
		request entities.ServiceCategory
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

func (h *ServiceCategoryHandler) GetAllServiceCategories(c echo.Context) error {
	var (
		svcName = "GetAllServiceCategories"
		request entities.DTRequest[entities.ServiceCategory]
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
		request entities.ServiceCategory
	)

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	request.Name = strings.ToUpper(request.Name)

	if err := h.serviceCategoryUsecase.UpdateServiceCategory(request); err != nil {
		utils.LoggMsg(svcName, "Failed to update service category", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update service category", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Service category updated successfully")
}

func (h *ServiceCategoryHandler) DeleteServiceCategory(c echo.Context) error {
	var (
		svcName = "DeleteServiceCategory"
		request entities.ServiceCategory
	)

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := h.serviceCategoryUsecase.DeleteServiceCategory(request.ID); err != nil {
		utils.LoggMsg(svcName, "Failed to delete service category", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete service category", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Service category deleted successfully")
}
