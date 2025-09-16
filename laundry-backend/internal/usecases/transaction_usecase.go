package usecases

import (
	"fmt"
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
)

type transactionUsecase struct {
	transactionRepo repositories.TransactionRepository
}

func NewTransactionUsecase(transactionRepo repositories.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: transactionRepo,
	}
}

func (u *transactionUsecase) GetAllTransactions() ([]entities.Transaction, error) {
	return u.transactionRepo.FindAll()
}

func (u *transactionUsecase) GetAllTransactionsDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error) {
	// Get order column
	var orderBy string
	var orderDir string
	if len(request.Order) > 0 && request.Order[0].Column < len(request.Columns) {
		orderBy = request.Columns[request.Order[0].Column].Data
		orderDir = request.Order[0].Dir
	}

	// Map column names to database column names
	columnMap := map[string]string{
		"id":              "id_transaksi",
		"nomor_invoice":   "nomor_invoice",
		"tanggal_masuk":   "tanggal_masuk",
		"status_transaksi": "status_transaksi",
		"total_harga":     "total_harga",
		"created_at":      "created_at",
	}

	if dbColumn, exists := columnMap[orderBy]; exists {
		orderBy = dbColumn
	} else {
		orderBy = "id_transaksi"
	}

	transactions, totalCount, err := u.transactionRepo.FindAllWithPagination(
		request.Length,
		request.Start,
		request.Search.Value,
		orderBy,
		orderDir,
	)
	if err != nil {
		return nil, err
	}

	response := &entities.DataTablesResponse{
		Draw:            request.Draw,
		RecordsTotal:    totalCount,
		RecordsFiltered: totalCount,
		Data:            transactions,
	}

	return response, nil
}

func (u *transactionUsecase) GetTransactionByID(id int) (*entities.Transaction, error) {
	return u.transactionRepo.FindByID(id)
}

func (u *transactionUsecase) GetTransactionsByOutletID(outletID int) ([]entities.Transaction, error) {
	return u.transactionRepo.FindByOutletID(outletID)
}

func (u *transactionUsecase) GetTransactionDetails(transactionID int) ([]entities.TransactionDetail, error) {
	return u.transactionRepo.FindDetailsByTransactionID(transactionID)
}

func (u *transactionUsecase) UpdateTransactionStatus(id int, request entities.UpdateTransactionStatusRequest) error {
	// Validate the status value
	validStatuses := map[string]bool{
		"diterima":  true,
		"diproses":  true,
		"selesai":   true,
		"diambil":   true,
	}
	
	if !validStatuses[request.Status] {
		return fmt.Errorf("invalid transaction status: %s", request.Status)
	}
	
	return u.transactionRepo.UpdateTransactionStatus(id, request.Status)
}

func (u *transactionUsecase) UpdatePaymentStatus(id int, request entities.UpdatePaymentStatusRequest) error {
	// Validate the status value
	validStatuses := map[string]bool{
		"lunas":       true,
		"belum lunas": true,
	}
	
	if !validStatuses[request.Status] {
		return fmt.Errorf("invalid payment status: %s", request.Status)
	}
	
	return u.transactionRepo.UpdatePaymentStatus(id, request.Status)
}

func (u *transactionUsecase) ProcessPaymentCallback(request entities.PaymentCallbackRequest) error {
	// Validate the payment status value
	validStatuses := map[string]bool{
		"lunas":       true,
		"belum lunas": true,
		"gagal":       true,
	}
	
	if !validStatuses[request.PaymentStatus] {
		return fmt.Errorf("invalid payment status: %s", request.PaymentStatus)
	}
	
	// Validate the payment method value
	validPaymentMethods := map[string]bool{
		"tunai":     true,
		"transfer":  true,
		"e-wallet":  true,
	}
	
	if request.PaymentMethod != "" && !validPaymentMethods[request.PaymentMethod] {
		return fmt.Errorf("invalid payment method: %s", request.PaymentMethod)
	}
	
	// First, check if the transaction exists
	transaction, err := u.transactionRepo.FindByID(request.TransactionID)
	if err != nil {
		return fmt.Errorf("failed to find transaction: %w", err)
	}
	
	if transaction == nil {
		return fmt.Errorf("transaction not found with id: %d", request.TransactionID)
	}
	
	return u.transactionRepo.UpdatePaymentCallback(request.TransactionID, request)
}