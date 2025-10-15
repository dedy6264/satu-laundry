package transactionrepo

import (
	"laundry-backend/entities"
)

type TransactionRepo interface {
	FindAll() (response []entities.Transaction, err error)
	FindAllWithPagination(request entities.DTRequest[entities.Transaction]) (response []entities.Transaction, totalCount int, err error)
	FindByID(id int) (response entities.Transaction, err error)
	FindByOutletID(outletID int) ([]entities.Transaction, error)
	FindDetailsByTransactionID(transactionID int) ([]entities.TransactionDetail, error)
	UpdateTransactionStatus(id int, status string) error
	UpdatePaymentStatus(id int, status string) error
	UpdatePaymentCallback(transactionID int, request entities.PaymentCallbackRequest) error
}
