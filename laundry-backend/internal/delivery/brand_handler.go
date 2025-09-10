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

type BrandHandler struct {
	brandUsecase usecases.BrandUsecase
}

func NewBrandHandler(brandUsecase usecases.BrandUsecase) *BrandHandler {
	return &BrandHandler{
		brandUsecase: brandUsecase,
	}
}

func (h *BrandHandler) CreateBrand(c echo.Context) error {
	var (
		request entities.RegisterBrandRequest
		svcName = "CreateBrand"
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	request.Name = strings.ToUpper(request.Name)
	request.PICName = strings.ToUpper(request.PICName)
	if err := h.brandUsecase.CreateBrand(request); err != nil {
		utils.LoggMsg(svcName, "Failed to create brand", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create brand", err.Error())
	}

	return MessageResponse(c, http.StatusCreated, "Brand created successfully")
}

func (h *BrandHandler) GetBrandByID(c echo.Context) error {
	var (
		svcName = "GetBrandByID"
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid brand ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid brand ID", err.Error())
	}

	brand, err := h.brandUsecase.GetBrandByID(id)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get brand", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get brand", err.Error())
	}

	if brand == nil {
		utils.LoggMsg(svcName, "Brand not found", nil)
		return ErrorResponse(c, http.StatusNotFound, "Brand not found", "Brand with given ID does not exist")
	}

	return SuccessResponse(c, http.StatusOK, "Brand retrieved successfully", brand)
}

func (h *BrandHandler) GetAllBrands(c echo.Context) error {
	var (
		svcName = "GetAllBrands"
	)
	brands, err := h.brandUsecase.GetAllBrands()
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get brands", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get brands", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Brands retrieved successfully", brands)
}

func (h *BrandHandler) UpdateBrand(c echo.Context) error {
	var (
		svcName = "UpdateBrand"
		request entities.RegisterBrandRequest
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid brand ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid brand ID", err.Error())
	}

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	request.Name = strings.ToUpper(request.Name)
	request.PICName = strings.ToUpper(request.PICName)
	if err := h.brandUsecase.UpdateBrand(id, request); err != nil {
		utils.LoggMsg(svcName, "Failed to update brand", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update brand", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Brand updated successfully")
}

func (h *BrandHandler) DeleteBrand(c echo.Context) error {
	var (
		svcName = "DeleteBrand"
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid brand ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid brand ID", err.Error())
	}

	if err := h.brandUsecase.DeleteBrand(id); err != nil {
		utils.LoggMsg(svcName, "Failed to delete brand", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete brand", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Brand deleted successfully")
}
