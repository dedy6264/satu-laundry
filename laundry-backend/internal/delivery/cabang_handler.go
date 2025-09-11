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

type CabangHandler struct {
	cabangUsecase usecases.CabangUsecase
}

func NewCabangHandler(cabangUsecase usecases.CabangUsecase) *CabangHandler {
	return &CabangHandler{
		cabangUsecase: cabangUsecase,
	}
}

func (h *CabangHandler) CreateCabang(c echo.Context) error {
	var (
		request entities.RegisterCabangRequest
		svcName = "CreateCabang"
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	request.Name = strings.ToUpper(request.Name)
	request.PICName = strings.ToUpper(request.PICName)
	if err := h.cabangUsecase.CreateCabang(request); err != nil {
		utils.LoggMsg(svcName, "Failed to create cabang", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create cabang", err.Error())
	}

	return MessageResponse(c, http.StatusCreated, "Cabang created successfully")
}

func (h *CabangHandler) GetCabangByID(c echo.Context) error {
	var (
		svcName = "GetCabangByID"
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid cabang ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid cabang ID", err.Error())
	}

	cabang, err := h.cabangUsecase.GetCabangByID(id)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get cabang", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get cabang", err.Error())
	}

	if cabang == nil {
		utils.LoggMsg(svcName, "Cabang not found", nil)
		return ErrorResponse(c, http.StatusNotFound, "Cabang not found", "Cabang with given ID does not exist")
	}

	return SuccessResponse(c, http.StatusOK, "Cabang retrieved successfully", cabang)
}

func (h *CabangHandler) GetCabangsByBrandID(c echo.Context) error {
	var (
		svcName = "GetCabangsByBrandID"
	)
	brandID, err := strconv.Atoi(c.Param("brand_id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid brand ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid brand ID", err.Error())
	}

	cabangs, err := h.cabangUsecase.GetCabangsByBrandID(brandID)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get cabangs", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get cabangs", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Cabangs retrieved successfully", cabangs)
}

func (h *CabangHandler) GetAllCabangs(c echo.Context) error {
	var (
		svcName = "GetAllCabangs"
		request entities.DataTablesRequest
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	response, err := h.cabangUsecase.GetAllCabangsDataTables(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get cabangs", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get cabangs", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Cabangs retrieved successfully", response)
}

func (h *CabangHandler) UpdateCabang(c echo.Context) error {
	var (
		svcName = "UpdateCabang"
		request entities.RegisterCabangRequest
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid cabang ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid cabang ID", err.Error())
	}

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	request.Name = strings.ToUpper(request.Name)
	request.PICName = strings.ToUpper(request.PICName)
	if err := h.cabangUsecase.UpdateCabang(id, request); err != nil {
		utils.LoggMsg(svcName, "Failed to update cabang", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update cabang", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Cabang updated successfully")
}

func (h *CabangHandler) DeleteCabang(c echo.Context) error {
	var (
		svcName = "DeleteCabang"
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid cabang ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid cabang ID", err.Error())
	}

	if err := h.cabangUsecase.DeleteCabang(id); err != nil {
		utils.LoggMsg(svcName, "Failed to delete cabang", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete cabang", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Cabang deleted successfully")
}
