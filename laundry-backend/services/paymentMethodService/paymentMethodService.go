package paymentmethodservice

import (
	"laundry-backend/entities"
	"laundry-backend/utils"

	"net/http"

	"github.com/labstack/echo/v4"
)

func (h PaymentMethodService) GetAllPaymentMethods(c echo.Context) error {
	var (
		svcName = "GetAllPaymentMethods"
		request entities.DTRequest[entities.PaymentMethod]
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	paymentMethods, totalCount, err := h.service.RepoPaymentMethod.FindAllWithPagination(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get payment methods", err)
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get payment methods", err.Error())
	}

	response := entities.DataTablesResponse{
		Draw:            request.Draw,
		RecordsTotal:    totalCount,
		RecordsFiltered: totalCount,
		Data:            paymentMethods,
	}
	return utils.SuccessResponse(c, http.StatusOK, "Payment methods retrieved successfully", response)
}
