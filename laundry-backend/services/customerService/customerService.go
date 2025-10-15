package customerservice

import (
	"database/sql"
	"laundry-backend/entities"
	"laundry-backend/utils"
	"strings"

	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (h CustomerService) GetAllCustomers(c echo.Context) error {
	var (
		svcName  = "GetAllCustomers"
		request  entities.DTRequest[entities.Customer]
		outletID int
	)
	// Ambil token dari context
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	request.UserID = int(claims["user_id"].(float64)) // JSON number → float64 → int)
	userAccess, err := h.service.RepoUserAccess.FindByID(request.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.LoggMsg(svcName, "invalid userAccess", err)
			return utils.ErrorResponse(c, http.StatusBadRequest, "invalid userAccess", err.Error())
		}

		utils.LoggMsg(svcName, "invalid userAccess", err)
		return utils.ErrorResponse(c, http.StatusBadRequest, "invalid userAccess", err.Error())
	}

	// // 2. get outlet
	if userAccess.ReferenceLevel != "cabang" {
		switch userAccess.ReferenceLevel {
		case "pegawai":

			employee, err := h.service.RepoEmployee.FindByID(userAccess.ReferenceID)
			if err != nil {
				utils.LoggMsg(svcName, "employeeRepo.FindByID", err)
				return utils.ErrorResponse(c, http.StatusBadRequest, "employeeRepo.FindByID", err.Error())
			}
			outletID = employee.OutletID
		default:
			utils.LoggMsg(svcName, "invalid reference level", err)
			return utils.ErrorResponse(c, http.StatusBadRequest, "invalid reference level", "")
		}
	}

	request.Data.OutletID = outletID
	if request.Data.OutletID == 0 {
		utils.LoggMsg(svcName, "invalid Outlet ID", err)
		return utils.ErrorResponse(c, http.StatusBadRequest, "invalid Outlet ID", "")
	}
	services, totalCount, err := h.service.RepoCustomer.FindAllWithPagination(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed", err)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Failed", err.Error())
	}

	response := entities.DataTablesResponse{
		Draw:            request.Draw,
		RecordsTotal:    totalCount,
		RecordsFiltered: totalCount,
		Data:            services,
	}

	return utils.SuccessResponse(c, http.StatusOK, "Services retrieved successfully", response)
}
func (h CustomerService) CreateCustomer(c echo.Context) error {
	var (
		svcName  = "CreateCustomer"
		request  entities.Customer
		outletID int
	)

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	if request.Name == "" {
		utils.LoggMsg(svcName, "Name is required", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Name is required", "")
	}

	if request.Phone == "" {
		utils.LoggMsg(svcName, "Phone is required", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Phone is required", "")
	}
	request.Name = strings.ToUpper(request.Name)
	UserID := int(claims["user_id"].(float64)) // JSON number → float64 → int)
	userAccess, err := h.service.RepoUserAccess.FindByID(UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.LoggMsg(svcName, "invalid userAccess", err)
			return utils.ErrorResponse(c, http.StatusBadRequest, "invalid userAccess", err.Error())
		}

		utils.LoggMsg(svcName, "invalid userAccess", err)
		return utils.ErrorResponse(c, http.StatusBadRequest, "invalid userAccess", err.Error())
	}

	// // 2. get outlet
	if userAccess.ReferenceLevel != "cabang" {
		switch userAccess.ReferenceLevel {
		case "pegawai":

			employee, err := h.service.RepoEmployee.FindByID(userAccess.ReferenceID)
			if err != nil {
				utils.LoggMsg(svcName, "employeeRepo.FindByID", err)
				return utils.ErrorResponse(c, http.StatusBadRequest, "employeeRepo.FindByID", err.Error())
			}
			outletID = employee.OutletID
		default:
			utils.LoggMsg(svcName, "invalid reference level", err)
			return utils.ErrorResponse(c, http.StatusBadRequest, "invalid reference level", "")
		}
	}
	request.OutletID = outletID
	if request.OutletID == 0 {
		utils.LoggMsg(svcName, "Outlet ID is required", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Outlet ID is required", "")
	}
	err = h.service.RepoCustomer.Create(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to create customer", err)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Failed to create customer", err.Error())
	}

	return utils.MessageResponse(c, http.StatusCreated, "Customer created successfully")
}
