package delivery

import (
	"net/http"
	"strconv"
	"laundry-backend/internal/entities"
	"laundry-backend/internal/usecases"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	customerUsecase usecases.CustomerUsecase
}

func NewCustomerHandler(customerUsecase usecases.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{
		customerUsecase: customerUsecase,
	}
}

func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	var request entities.RegisterCustomerRequest
	if err := c.Bind(&request); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	// Validate required fields
	if request.Name == "" {
		return ErrorResponse(c, http.StatusBadRequest, "Name is required", "")
	}

	err := h.customerUsecase.CreateCustomer(request)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Failed to create customer", err.Error())
	}

	return MessageResponse(c, http.StatusCreated, "Customer created successfully")
}

func (h *CustomerHandler) GetCustomerByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid customer ID", err.Error())
	}

	customer, err := h.customerUsecase.GetCustomerByID(id)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get customer", err.Error())
	}

	if customer == nil {
		return ErrorResponse(c, http.StatusNotFound, "Customer not found", "")
	}

	return SuccessResponse(c, http.StatusOK, "Customer retrieved successfully", customer)
}

func (h *CustomerHandler) GetCustomersByOutletID(c echo.Context) error {
	outletID, err := strconv.Atoi(c.Param("outlet_id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid outlet ID", err.Error())
	}

	customers, err := h.customerUsecase.GetCustomersByOutletID(outletID)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get customers", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Customers retrieved successfully", customers)
}

func (h *CustomerHandler) GetAllCustomers(c echo.Context) error {
	var request entities.DataTablesRequest
	if err := c.Bind(&request); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	response, err := h.customerUsecase.GetAllCustomersDataTables(request)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get customers", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Customers retrieved successfully", response)
}

func (h *CustomerHandler) GetAllCustomersDataTables(c echo.Context) error {
	var request entities.DataTablesRequest
	if err := c.Bind(&request); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	response, err := h.customerUsecase.GetAllCustomersDataTables(request)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get customers", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Customers retrieved successfully", response)
}

func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid customer ID", err.Error())
	}

	var request entities.RegisterCustomerRequest
	if err := c.Bind(&request); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	err = h.customerUsecase.UpdateCustomer(id, request)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Failed to update customer", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Customer updated successfully")
}

func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid customer ID", err.Error())
	}

	err = h.customerUsecase.DeleteCustomer(id)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Failed to delete customer", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Customer deleted successfully")
}