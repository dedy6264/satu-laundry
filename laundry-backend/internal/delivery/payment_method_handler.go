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

type PaymentMethodHandler struct {
	paymentMethodUsecase usecases.PaymentMethodUsecase
}

func NewPaymentMethodHandler(paymentMethodUsecase usecases.PaymentMethodUsecase) *PaymentMethodHandler {
	return &PaymentMethodHandler{
		paymentMethodUsecase: paymentMethodUsecase,
	}
}

func (h *PaymentMethodHandler) CreatePaymentMethod(c echo.Context) error {
	var (
		request entities.CreatePaymentMethodRequest
		svcName = "CreatePaymentMethod"
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	request.NamaMetode = strings.ToUpper(request.NamaMetode)
	if err := h.paymentMethodUsecase.CreatePaymentMethod(request); err != nil {
		utils.LoggMsg(svcName, "Failed to create payment method", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create payment method", err.Error())
	}

	return MessageResponse(c, http.StatusCreated, "Payment method created successfully")
}

func (h *PaymentMethodHandler) GetPaymentMethodByID(c echo.Context) error {
	var (
		svcName = "GetPaymentMethodByID"
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid payment method ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid payment method ID", err.Error())
	}

	paymentMethod, err := h.paymentMethodUsecase.GetPaymentMethodByID(id)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get payment method", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get payment method", err.Error())
	}

	if paymentMethod == nil {
		utils.LoggMsg(svcName, "Payment method not found", nil)
		return ErrorResponse(c, http.StatusNotFound, "Payment method not found", "Payment method with given ID does not exist")
	}

	return SuccessResponse(c, http.StatusOK, "Payment method retrieved successfully", paymentMethod)
}

func (h *PaymentMethodHandler) GetAllPaymentMethods(c echo.Context) error {
	var (
		svcName = "GetAllPaymentMethods"
		request entities.DataTablesRequest
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	response, err := h.paymentMethodUsecase.GetAllPaymentMethodsDataTables(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get payment methods", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get payment methods", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Payment methods retrieved successfully", response)
}

func (h *PaymentMethodHandler) UpdatePaymentMethod(c echo.Context) error {
	var (
		svcName = "UpdatePaymentMethod"
		request entities.UpdatePaymentMethodRequest
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid payment method ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid payment method ID", err.Error())
	}

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Invalid request format", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	request.NamaMetode = strings.ToUpper(request.NamaMetode)
	if err := h.paymentMethodUsecase.UpdatePaymentMethod(id, request); err != nil {
		utils.LoggMsg(svcName, "Failed to update payment method", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update payment method", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Payment method updated successfully")
}

func (h *PaymentMethodHandler) DeletePaymentMethod(c echo.Context) error {
	var (
		svcName = "DeletePaymentMethod"
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid payment method ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid payment method ID", err.Error())
	}

	if err := h.paymentMethodUsecase.DeletePaymentMethod(id); err != nil {
		utils.LoggMsg(svcName, "Failed to delete payment method", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete payment method", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Payment method deleted successfully")
}