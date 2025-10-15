package useraccessservice

import (
	"laundry-backend/constant"
	"laundry-backend/entities"
	"laundry-backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h UserAccessService) UserLogin(c echo.Context) error {
	var (
		request  entities.UserLoginRequest
		svcName  = "UserLogin"
		response entities.UserLoginResponse
		id       int
	)

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return utils.ErrorResponse(c, http.StatusBadRequest, constant.InvalidParameterMessage, err.Error())
	}
	if request.Username == "" || request.Password == "" {
		utils.LoggMsg(svcName, "Username and password are required", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, constant.InvalidParameterMessage, "Username and password are required")
	}

	userAccess, err := h.service.RepoUserAccess.AuthenticateUser(request.Username, request.Password)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to authenticate user", err)
		return utils.ErrorResponse(c, http.StatusInternalServerError, constant.InvalidParameterMessage, "Failed to authenticate user")
	}
	//get data hirarki by reference
	switch userAccess.ReferenceLevel {
	case "cabang":
		cabang, err := h.service.RepoCabang.FindByID(userAccess.ReferenceID)
		if err != nil {
			return utils.ErrorResponse(c, http.StatusBadRequest, constant.NotFoundMessage, "")
		}
		id = cabang.ID
	case "outlet":
		outlet, err := h.service.RepoOutlet.FindByID(userAccess.ReferenceID)
		if err != nil {
			return utils.ErrorResponse(c, http.StatusBadRequest, constant.NotFoundMessage, "")
		}
		id = outlet.ID
	default: //karyawan
		employee, err := h.service.RepoEmployee.FindByID(userAccess.ReferenceID)
		if err != nil {
			return utils.ErrorResponse(c, http.StatusBadRequest, constant.NotFoundMessage, "")
		}
		id = employee.ID
	}
	// Generate JWT token
	tokenString, err := utils.GenerateJWT(userAccess.ID, id, userAccess.Username, userAccess.Role, userAccess.ReferenceLevel)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, constant.InternalErrorMessage, err.Error())
	}

	response = entities.UserLoginResponse{
		AccessToken: tokenString,
		User:        userAccess,
	}

	return utils.SuccessResponse(c, http.StatusOK, "Login successful", response)
}
