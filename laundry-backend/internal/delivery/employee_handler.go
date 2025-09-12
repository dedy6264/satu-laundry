package delivery

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EmployeeHandler struct {
	employeeUsecase usecases.EmployeeUsecase
}

func NewEmployeeHandler(employeeUsecase usecases.EmployeeUsecase) *EmployeeHandler {
	return &EmployeeHandler{
		employeeUsecase: employeeUsecase,
	}
}

func (h *EmployeeHandler) CreateEmployee(c echo.Context) error {
	var request entities.RegisterEmployeeRequest
	if err := c.Bind(&request); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	// Validate required fields
	if request.NIK == "" || request.Name == "" {
		return ErrorResponse(c, http.StatusBadRequest, "NIK and name are required", "")
	}

	err := h.employeeUsecase.CreateEmployee(request)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Failed to create employee", err.Error())
	}

	return MessageResponse(c, http.StatusCreated, "Employee created successfully")
}

func (h *EmployeeHandler) GetEmployeeByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid employee ID", err.Error())
	}

	employee, err := h.employeeUsecase.GetEmployeeByID(id)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get employee", err.Error())
	}

	if employee == nil {
		return ErrorResponse(c, http.StatusNotFound, "Employee not found", "")
	}

	return SuccessResponse(c, http.StatusOK, "Employee retrieved successfully", employee)
}

func (h *EmployeeHandler) GetAllEmployees(c echo.Context) error {
	var request entities.DataTablesRequest
	if err := c.Bind(&request); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	response, err := h.employeeUsecase.GetAllEmployeesDataTables(request)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get employees", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Employees retrieved successfully", response)
}

func (h *EmployeeHandler) GetAllEmployeesDataTables(c echo.Context) error {
	var request entities.DataTablesRequest
	if err := c.Bind(&request); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	response, err := h.employeeUsecase.GetAllEmployeesDataTables(request)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get employees", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Employees retrieved successfully", response)
}

func (h *EmployeeHandler) UpdateEmployee(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid employee ID", err.Error())
	}

	var request entities.RegisterEmployeeRequest
	if err := c.Bind(&request); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	err = h.employeeUsecase.UpdateEmployee(id, request)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Failed to update employee", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Employee updated successfully")
}

func (h *EmployeeHandler) DeleteEmployee(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid employee ID", err.Error())
	}

	err = h.employeeUsecase.DeleteEmployee(id)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Failed to delete employee", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Employee deleted successfully")
}
