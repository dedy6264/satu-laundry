package usecases

import (
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