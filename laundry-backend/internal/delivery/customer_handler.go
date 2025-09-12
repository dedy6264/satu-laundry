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

type CustomerHandler struct {
	customerUsecase usecases.CustomerUsecase
}

func NewCustomerHandler(customerUsecase usecases.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{
		customerUsecase: customerUsecase,
	}
}

func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	var (
		svcName = "CreateCustomer"
		request entities.RegisterCustomerRequest
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	// Validate required fields
	if request.Name == "" {
		utils.LoggMsg(svcName, "Name is required", nil)
		return ErrorResponse(c, http.StatusBadRequest, "Name is required", "")
	}
	request.Name = strings.ToUpper(request.Name)
	err := h.customerUsecase.CreateCustomer(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to create customer", err)
		return ErrorResponse(c, http.StatusBadRequest, "Failed to create customer", err.Error())
	}

	return MessageResponse(c, http.StatusCreated, "Customer created successfully")
}

func (h *CustomerHandler) GetCustomerByID(c echo.Context) error {
	var (
		svcName = "GetCustomerByID"
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid customer ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid customer ID", err.Error())
	}

	customer, err := h.customerUsecase.GetCustomerByID(id)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get customer", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get customer", err.Error())
	}

	if customer == nil {
		utils.LoggMsg(svcName, "Customer not found", nil)
		return ErrorResponse(c, http.StatusNotFound, "Customer not found", "")
	}

	return SuccessResponse(c, http.StatusOK, "Customer retrieved successfully", customer)
}

func (h *CustomerHandler) GetCustomersByOutletID(c echo.Context) error {
	var (
		svcName = "GetCustomersByOutletID"
	)
	outletID, err := strconv.Atoi(c.Param("outlet_id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid outlet ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid outlet ID", err.Error())
	}

	customers, err := h.customerUsecase.GetCustomersByOutletID(outletID)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get customersID", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get customers", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Customers retrieved successfully", customers)
}

func (h *CustomerHandler) GetAllCustomers(c echo.Context) error {
	var (
		svcName = "GetAllCustomers"
		request entities.DataTablesRequest
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	response, err := h.customerUsecase.GetAllCustomersDataTables(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get customers", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get customers", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Customers retrieved successfully", response)
}

func (h *CustomerHandler) GetAllCustomersDataTables(c echo.Context) error {
	var (
		svcName = "GetAllCustomersDataTables"
		request entities.DataTablesRequest
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	response, err := h.customerUsecase.GetAllCustomersDataTables(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get customers", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get customers", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Customers retrieved successfully", response)
}

func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
	var (
		svcName = "UpdateCustomer"
		request entities.RegisterCustomerRequest
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid customer ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid customer ID", err.Error())
	}

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	request.Name = strings.ToUpper(request.Name)
	err = h.customerUsecase.UpdateCustomer(id, request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to update customer", err)
		return ErrorResponse(c, http.StatusBadRequest, "Failed to update customer", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Customer updated successfully")
}

func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
	var (
		svcName = "DeleteCustomer"
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid customer ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid customer ID", err.Error())
	}

	err = h.customerUsecase.DeleteCustomer(id)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to delete customer", err)
		return ErrorResponse(c, http.StatusBadRequest, "Failed to delete customer", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Customer deleted successfully")
}
