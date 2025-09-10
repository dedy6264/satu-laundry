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

type OutletHandler struct {
	outletUsecase usecases.OutletUsecase
}

func NewOutletHandler(outletUsecase usecases.OutletUsecase) *OutletHandler {
	return &OutletHandler{
		outletUsecase: outletUsecase,
	}
}

func (h *OutletHandler) CreateOutlet(c echo.Context) error {
	var (
		request entities.RegisterOutletRequest
		svcName = "CreateOutlet"
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	request.Name = strings.ToUpper(request.Name)
	request.PICName = strings.ToUpper(request.PICName)
	if err := h.outletUsecase.CreateOutlet(request); err != nil {
		utils.LoggMsg(svcName, "Failed to create outlet", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create outlet", err.Error())
	}

	return MessageResponse(c, http.StatusCreated, "Outlet created successfully")
}

func (h *OutletHandler) GetOutletByID(c echo.Context) error {
	var (
		svcName = "GetOutletByID"
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid outlet ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid outlet ID", err.Error())
	}

	outlet, err := h.outletUsecase.GetOutletByID(id)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get outlet", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get outlet", err.Error())
	}

	if outlet == nil {
		utils.LoggMsg(svcName, "Outlet not found", nil)
		return ErrorResponse(c, http.StatusNotFound, "Outlet not found", "Outlet with given ID does not exist")
	}

	return SuccessResponse(c, http.StatusOK, "Outlet retrieved successfully", outlet)
}

func (h *OutletHandler) GetOutletsByCabangID(c echo.Context) error {
	var (
		svcName = "GetOutletsByCabangID"
	)
	cabangID, err := strconv.Atoi(c.Param("cabang_id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid cabang ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid cabang ID", err.Error())
	}

	outlets, err := h.outletUsecase.GetOutletsByCabangID(cabangID)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get outlets", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get outlets", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Outlets retrieved successfully", outlets)
}

func (h *OutletHandler) GetAllOutlets(c echo.Context) error {
	var (
		svcName = "GetAllOutlets"
	)
	outlets, err := h.outletUsecase.GetAllOutlets()
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get outlets", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get outlets", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Outlets retrieved successfully", outlets)
}

func (h *OutletHandler) UpdateOutlet(c echo.Context) error {
	var (
		svcName = "UpdateOutlet"
		request entities.RegisterOutletRequest
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid outlet ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid outlet ID", err.Error())
	}

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	request.Name = strings.ToUpper(request.Name)
	request.PICName = strings.ToUpper(request.PICName)
	if err := h.outletUsecase.UpdateOutlet(id, request); err != nil {
		utils.LoggMsg(svcName, "Failed to update outlet", nil)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update outlet", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Outlet updated successfully")
}

func (h *OutletHandler) DeleteOutlet(c echo.Context) error {
	var (
		svcName = "DeleteOutlet"
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid outlet ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid outlet ID", err.Error())
	}

	if err := h.outletUsecase.DeleteOutlet(id); err != nil {
		utils.LoggMsg(svcName, "Failed to delete outlet", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete outlet", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Outlet deleted successfully")
}
