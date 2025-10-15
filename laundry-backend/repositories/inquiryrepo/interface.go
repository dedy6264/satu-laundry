package inquiryrepo

import (
	"database/sql"
	"laundry-backend/entities"
)

type InquiryRepo interface {
	ValidateCustomer(id int) (bool, error)
	BeginTransaction() (*sql.Tx, error)
	InsertTransactionWithTx(tx *sql.Tx, transaction entities.Transaction) (int, error)
	InsertTransactionDetailWithTx(tx *sql.Tx, detail []entities.TransactionDetail) (err error)
	InsertPaymentWithTx(tx *sql.Tx, payment entities.Payment) error
	InsertHistoryStatusTransactionWithTx(tx *sql.Tx, history entities.HistoryStatusTransaction) error
}
