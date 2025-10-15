package serviceservice

import (
	"database/sql"
	"laundry-backend/entities"
	"laundry-backend/utils"

	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (h ServiceService) GetAllServices(c echo.Context) error {
	var (
		svcName = "GetAllServices"
		request entities.DTRequest[entities.Service]
		brandID int
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
			brandID = employee.BrandID
		case "outlet":
			outlet, err := h.service.RepoOutlet.FindByID(userAccess.ReferenceID)
			if err != nil {
				utils.LoggMsg(svcName, "outletRepo.FindByID", err)
				return utils.ErrorResponse(c, http.StatusBadRequest, "outletRepo.FindByID", err.Error())
			}
			brandID = outlet.BrandID
		default:
			utils.LoggMsg(svcName, "invalid reference level", err)
			return utils.ErrorResponse(c, http.StatusBadRequest, "invalid reference level", err.Error())
		}
	}

	request.Data.BrandID = brandID
	services, totalCount, err := h.service.RepoService.FindAllWithPagination(request)
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
