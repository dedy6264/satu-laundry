package delivery

import (
	"fmt"
	"laundry-backend/internal/entities"
	"laundry-backend/internal/usecases"
	"laundry-backend/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EmployeeAccessHandler struct {
	employeeAccessUsecase usecases.EmployeeAccessUsecase
	jwtSecret             string
}

func NewEmployeeAccessHandler(
	employeeAccessUsecase usecases.EmployeeAccessUsecase,
	jwtSecret string,
) *EmployeeAccessHandler {
	return &EmployeeAccessHandler{
		employeeAccessUsecase: employeeAccessUsecase,
		jwtSecret:             jwtSecret,
	}
}

func (h *EmployeeAccessHandler) CreateEmployeeAccess(c echo.Context) error {
	var (
		request entities.CreateEmployeeAccessRequest
		svcName = "CreateEmployeeAccess"
	)

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := h.employeeAccessUsecase.CreateEmployeeAccess(request); err != nil {
		utils.LoggMsg(svcName, "Failed to create employee access", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create employee access", err.Error())
	}

	return MessageResponse(c, http.StatusCreated, "Employee access created successfully")
}

func (h *EmployeeAccessHandler) GetEmployeeAccessByID(c echo.Context) error {
	var (
		svcName = "GetEmployeeAccessByID"
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid employee access ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid employee access ID", err.Error())
	}

	access, err := h.employeeAccessUsecase.GetEmployeeAccessByID(id)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get employee access", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get employee access", err.Error())
	}

	if access == nil {
		return ErrorResponse(c, http.StatusNotFound, "Employee access not found", "")
	}

	return SuccessResponse(c, http.StatusOK, "Employee access retrieved successfully", access)
}

func (h *EmployeeAccessHandler) GetAllEmployeeAccess(c echo.Context) error {
	var (
		svcName = "GetAllEmployeeAccess"
	)

	accesses, err := h.employeeAccessUsecase.GetAllEmployeeAccess()
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get all employee access", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get all employee access", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Employee accesses retrieved successfully", accesses)
}

func (h *EmployeeAccessHandler) GetAllEmployeeAccessDataTables(c echo.Context) error {
	var (
		request entities.DataTablesRequest
		svcName = "GetAllEmployeeAccessDataTables"
	)

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	response, err := h.employeeAccessUsecase.GetAllEmployeeAccessDataTables(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get employee access data tables", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get employee access data tables", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Employee accesses retrieved successfully", response)
}

func (h *EmployeeAccessHandler) UpdateEmployeeAccess(c echo.Context) error {
	var (
		request entities.UpdateEmployeeAccessRequest
		svcName = "UpdateEmployeeAccess"
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid employee access ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid employee access ID", err.Error())
	}

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := h.employeeAccessUsecase.UpdateEmployeeAccess(id, request); err != nil {
		utils.LoggMsg(svcName, "Failed to update employee access", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update employee access", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Employee access updated successfully")
}

func (h *EmployeeAccessHandler) UpdateEmployeePassword(c echo.Context) error {
	var (
		request entities.UpdateEmployeePasswordRequest
		svcName = "UpdateEmployeePassword"
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid employee access ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid employee access ID", err.Error())
	}

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := h.employeeAccessUsecase.UpdateEmployeePassword(id, request); err != nil {
		utils.LoggMsg(svcName, "Failed to update employee password", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update employee password", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Employee password updated successfully")
}

func (h *EmployeeAccessHandler) DeleteEmployeeAccess(c echo.Context) error {
	var (
		svcName = "DeleteEmployeeAccess"
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid employee access ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid employee access ID", err.Error())
	}

	if err := h.employeeAccessUsecase.DeleteEmployeeAccess(id); err != nil {
		utils.LoggMsg(svcName, "Failed to delete employee access", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete employee access", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Employee access deleted successfully")
}

func (h *EmployeeAccessHandler) GetEmployeeAccessByOutletID(c echo.Context) error {
	var (
		svcName = "GetEmployeeAccessByOutletID"
	)

	outletID, err := strconv.Atoi(c.Param("outlet_id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid outlet ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid outlet ID", err.Error())
	}

	accesses, err := h.employeeAccessUsecase.GetEmployeeAccessByOutletID(outletID)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get employee access by outlet ID", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get employee access by outlet ID", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Employee accesses retrieved successfully", accesses)
}

func (h *EmployeeAccessHandler) EmployeeLogin(c echo.Context) error {
	var (
		request entities.EmployeeLoginRequest
		svcName = "EmployeeLogin"
	)

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	if request.Username == "" || request.Password == "" {
		utils.LoggMsg(svcName, "Username and password are required", nil)
		return ErrorResponse(c, http.StatusBadRequest, "Username and password are required", "")
	}
	fmt.Println("::::::", request)
	response, err := h.employeeAccessUsecase.AuthenticateEmployee(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to authenticate employee", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to authenticate employee", err.Error())
	}

	if response == nil {
		return ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials", "")
	}

	return SuccessResponse(c, http.StatusOK, "Login successful", response)
}
