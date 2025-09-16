package delivery

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/usecases"
	"laundry-backend/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionUsecase usecases.TransactionUsecase
}

func NewTransactionHandler(transactionUsecase usecases.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		transactionUsecase: transactionUsecase,
	}
}

func (h *TransactionHandler) GetAllTransactions(c echo.Context) error {
	var (
		svcName = "GetAllTransactions"
		request entities.DataTablesRequest
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	response, err := h.transactionUsecase.GetAllTransactionsDataTables(request)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get transactions", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get transactions", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Transactions retrieved successfully", response)
}

func (h *TransactionHandler) GetTransactionByID(c echo.Context) error {
	var (
		svcName = "GetTransactionByID"
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid transaction ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid transaction ID", err.Error())
	}

	transaction, err := h.transactionUsecase.GetTransactionByID(id)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get transaction", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get transaction", err.Error())
	}

	if transaction == nil {
		utils.LoggMsg(svcName, "Transaction not found", nil)
		return ErrorResponse(c, http.StatusNotFound, "Transaction not found", "Transaction with given ID does not exist")
	}

	return SuccessResponse(c, http.StatusOK, "Transaction retrieved successfully", transaction)
}

func (h *TransactionHandler) GetTransactionsByOutletID(c echo.Context) error {
	var (
		svcName = "GetTransactionsByOutletID"
	)
	outletID, err := strconv.Atoi(c.Param("outlet_id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid outlet ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid outlet ID", err.Error())
	}

	transactions, err := h.transactionUsecase.GetTransactionsByOutletID(outletID)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get transactions by outlet", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get transactions by outlet", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Transactions retrieved successfully", transactions)
}

func (h *TransactionHandler) GetTransactionDetails(c echo.Context) error {
	var (
		svcName = "GetTransactionDetails"
	)
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid transaction ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid transaction ID", err.Error())
	}

	details, err := h.transactionUsecase.GetTransactionDetails(transactionID)
	if err != nil {
		utils.LoggMsg(svcName, "Failed to get transaction details", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to get transaction details", err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Transaction details retrieved successfully", details)
}

func (h *TransactionHandler) UpdateTransactionStatus(c echo.Context) error {
	var (
		svcName = "UpdateTransactionStatus"
		request entities.UpdateTransactionStatusRequest
	)
	
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid transaction ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid transaction ID", err.Error())
	}

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := h.transactionUsecase.UpdateTransactionStatus(transactionID, request); err != nil {
		utils.LoggMsg(svcName, "Failed to update transaction status", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update transaction status", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Transaction status updated successfully")
}

func (h *TransactionHandler) UpdatePaymentStatus(c echo.Context) error {
	var (
		svcName = "UpdatePaymentStatus"
		request entities.UpdatePaymentStatusRequest
	)
	
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LoggMsg(svcName, "Invalid transaction ID", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid transaction ID", err.Error())
	}

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := h.transactionUsecase.UpdatePaymentStatus(transactionID, request); err != nil {
		utils.LoggMsg(svcName, "Failed to update payment status", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update payment status", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Payment status updated successfully")
}

func (h *TransactionHandler) ProcessPaymentCallback(c echo.Context) error {
	var (
		svcName = "ProcessPaymentCallback"
		request entities.PaymentCallbackRequest
	)

	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := h.transactionUsecase.ProcessPaymentCallback(request); err != nil {
		utils.LoggMsg(svcName, "Failed to process payment callback", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to process payment callback", err.Error())
	}

	return MessageResponse(c, http.StatusOK, "Payment callback processed successfully")
}