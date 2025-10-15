package delivery

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/usecases"
	"laundry-backend/internal/utils"
	"net/http"

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

func (h *PaymentMethodHandler) GetAllPaymentMethods(c echo.Context) error {
	var (
		svcName = "GetAllPaymentMethods"
		request entities.DTRequest[entities.PaymentMethod]
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
