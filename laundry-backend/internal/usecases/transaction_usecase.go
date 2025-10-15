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

func (u *transactionUsecase) GetAllTransactionsDataTables(request entities.DTRequest[entities.Transaction]) (*entities.DataTablesResponse, error) {

	// Get order column

	transactions, totalCount, err := u.transactionRepo.FindAllWithPagination(request)
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

func (u *transactionUsecase) GetTransactionByID(id int) (entities.Transaction, error) {
	return u.transactionRepo.FindByID(id)
}

func (u *transactionUsecase) GetTransactionsByOutletID(outletID int) ([]entities.Transaction, error) {
	return u.transactionRepo.FindByOutletID(outletID)
}

func (u *transactionUsecase) GetTransactionDetails(request entities.Transaction) (result entities.TransactionComplete, err error) {
	details, err := u.transactionRepo.FindDetailsByTransactionID(request.ID)

	if err != nil {
		return result, err
	}
	trx, err := u.transactionRepo.FindByID(request.ID)
	if err != nil {
		return result, err
	}

	result = entities.TransactionComplete{
		ID:             trx.ID,
		CustomerID:     trx.CustomerID,
		CustomerName:   trx.CustomerName,
		OutletID:       trx.OutletID,
		OutletName:     trx.OutletName,
		UserID:         trx.UserID,
		InvoiceNumber:  trx.InvoiceNumber,
		EntryDate:      trx.EntryDate,
		CompletionDate: trx.CompletionDate,
		PickupDate:     trx.PickupDate,
		TotalPrice:     trx.TotalPrice,
		PaidAmount:     trx.PaidAmount,
		ChangeAmount:   trx.ChangeAmount,
		Status:         trx.Status,
		Note:           trx.Note,
		Detail:         details,
	}
	return result, nil
}

func (u *transactionUsecase) UpdateTransactionStatus(id int, request entities.UpdateTransactionStatusRequest) error {
	// Validate the status value
	validStatuses := map[string]bool{
		"diterima": true,
		"diproses": true,
		"selesai":  true,
		"diambil":  true,
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
		"tunai":    true,
		"transfer": true,
		"e-wallet": true,
	}

	if request.PaymentMethod != "" && !validPaymentMethods[request.PaymentMethod] {
		return fmt.Errorf("invalid payment method: %s", request.PaymentMethod)
	}

	// First, check if the transaction exists
	_, err := u.transactionRepo.FindByID(request.TransactionID)
	if err != nil {
		return fmt.Errorf("failed to find transaction: %w", err)
	}

	// if transaction == nil {
	// 	return fmt.Errorf("transaction not found with id: %d", request.TransactionID)
	// }

	return u.transactionRepo.UpdatePaymentCallback(request.TransactionID, request)
}
