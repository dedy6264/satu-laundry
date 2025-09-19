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

type UserAccessHandler struct {
	userAccessUsecase usecases.UserAccessUsecase
	jwtSecret         string
}

func NewUserAccessHandler(
	userAccessUsecase usecases.UserAccessUsecase,
	jwtSecret string,
) *UserAccessHandler {
	return &UserAccessHandler{
		userAccessUsecase: userAccessUsecase,
		jwtSecret:         jwtSecret,
	}
}

func (h *UserAccessHandler) CreateUserAccess(c echo.Context) error {
	var (
		request entities.CreateUserAccessRequest
		svcName = "CreateUserAccess"
	)

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := h.userAccessUsecase.CreateUserAccess(request); err != nil {
		utils.LoggMsg(svcName, "Failed to create user access", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create user access", err.Error())
	}

	return MessageResponse(c, http.StatusCreated, "User access created successfully")
}

func (h *UserAccessHandler) GetUserAccessByID(c echo.Context) error {
	var (
		svcName = "GetUserAccessByID"
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid user access ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid user access ID", err.Error())
	}

	access, err := h.userAccessUsecase.GetUserAccessByID(id)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get user access", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get user access", err.Error())
	}

	if access == nil {
		return ErrorResponse(c, http.StatusNotFound, "User access not found", "")
	}

	return SuccessResponse(c, http.StatusOK, "User access retrieved successfully", access)
}

func (h *UserAccessHandler) GetAllUserAccess(c echo.Context) error {
	var (
		svcName = "GetAllUserAccess"
	)

	accesses, err := h.userAccessUsecase.GetAllUserAccess()
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get all user access", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get all user access", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "User accesses retrieved successfully", accesses)
}

func (h *UserAccessHandler) GetAllUserAccessDataTables(c echo.Context) error {
	var (
		request entities.DataTablesRequest
		svcName = "GetAllUserAccessDataTables"
	)

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	response, err := h.userAccessUsecase.GetAllUserAccessDataTables(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get user access data tables", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get user access data tables", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "User accesses retrieved successfully", response)
}

func (h *UserAccessHandler) UpdateUserAccess(c echo.Context) error {
	var (
		request entities.UpdateUserAccessRequest
		svcName = "UpdateUserAccess"
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid user access ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid user access ID", err.Error())
	}

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := h.userAccessUsecase.UpdateUserAccess(id, request); err != nil {
		utils.LoggMsg(svcName, "Failed to update user access", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update user access", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "User access updated successfully")
}

func (h *UserAccessHandler) UpdateUserPassword(c echo.Context) error {
	var (
		request entities.UpdateUserPasswordRequest
		svcName = "UpdateUserPassword"
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid user access ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid user access ID", err.Error())
	}

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := h.userAccessUsecase.UpdateUserPassword(id, request); err != nil {
		utils.LoggMsg(svcName, "Failed to update user password", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update user password", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "User password updated successfully")
}

func (h *UserAccessHandler) DeleteUserAccess(c echo.Context) error {
	var (
		svcName = "DeleteUserAccess"
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid user access ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid user access ID", err.Error())
	}

	if err := h.userAccessUsecase.DeleteUserAccess(id); err != nil {
		utils.LoggMsg(svcName, "Failed to delete user access", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete user access", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "User access deleted successfully")
}

func (h *UserAccessHandler) UserLogin(c echo.Context) error {
	var (
		request entities.UserLoginRequest
		svcName = "UserLogin"
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
	response, err := h.userAccessUsecase.AuthenticateUser(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to authenticate user", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to authenticate user", err.Error())
	}

	if response == nil {
		return ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials", "")
	}

	return SuccessResponse(c, http.StatusOK, "Login successful", response)
}
