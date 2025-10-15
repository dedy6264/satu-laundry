package transactionservice

import (
	"laundry-backend/entities"
	"laundry-backend/utils"

	"net/http"

	"github.com/labstack/echo/v4"
)

func (h TransactionService) GetAllTransactions(c echo.Context) error {
	var (
		svcName = "GetAllTransactions"
		request entities.DTRequest[entities.Transaction]
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	transactions, totalCount, err := h.service.RepoTransaction.FindAllWithPagination(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get transactions", err)
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get transactions", err.Error())
	}

	response := entities.DataTablesResponse{
		Draw:            request.Draw,
		RecordsTotal:    totalCount,
		RecordsFiltered: totalCount,
		Data:            transactions,
	}
	return utils.SuccessResponse(c, http.StatusOK, "Transactions retrieved successfully", response)
}
