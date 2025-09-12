package repositories

import (
	"database/sql"
	"fmt"
	"laundry-backend/internal/entities"
	"time"
)

type inquiryPostgresRepository struct {
	db           *sql.DB
	employeeRepo EmployeeRepository
}

func NewInquiryRepository(db *sql.DB, employeeRepo EmployeeRepository) InquiryRepository {
	return &inquiryPostgresRepository{
		db:           db,
		employeeRepo: employeeRepo,
	}
}

func (r *inquiryPostgresRepository) ValidateServicePackage(id int) (bool, error) {
	query := `SELECT COUNT(*) FROM paket_layanan WHERE id_layanan = $1`
	row := r.db.QueryRow(query, id)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *inquiryPostgresRepository) ValidateEmployee(id int) (bool, error) {
	// Use the employee repository to find the employee by ID
	employee, err := r.employeeRepo.FindByID(id)
	if err != nil {
		return false, err
	}

	// If employee is not found, return false
	if employee == nil {
		return false, nil
	}

	// Employee exists, return true
	return true, nil
}

func (r *inquiryPostgresRepository) ValidateCustomer(id int) (bool, error) {
	query := `SELECT COUNT(*) FROM pelanggan WHERE id_pelanggan = $1`
	row := r.db.QueryRow(query, id)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *inquiryPostgresRepository) InsertTransaction(transaction *entities.Transaction) error {
	query := `INSERT INTO transaksi (
		id_pelanggan, id_outlet, nomor_invoice, tanggal_masuk, status, catatan
	) VALUES (
		$1, $2, $3, $4, $5, $6
	) RETURNING id_transaksi`

	var id int
	err := r.db.QueryRow(
		query,
		transaction.CustomerID,
		transaction.OutletID,
		transaction.InvoiceNumber,
		transaction.EntryDate,
		transaction.Status,
		transaction.Note,
	).Scan(&id)

	if err != nil {
		return err
	}

	transaction.ID = id
	// Set default values for timestamps
	now := time.Now()
	transaction.CreatedAt = now
	transaction.UpdatedAt = now

	return nil
}

func (r *inquiryPostgresRepository) InsertTransactionDetail(detail *entities.TransactionDetail) error {
	query := `INSERT INTO detail_transaksi (
		id_transaksi, id_layanan, jumlah, harga, subtotal
	) VALUES (
		$1, $2, $3, $4, $5
	) RETURNING id_detail`

	var id int
	err := r.db.QueryRow(
		query,
		detail.TransactionID,
		detail.ServiceID,
		detail.Quantity,
		detail.Price,
		detail.Subtotal,
	).Scan(&id)

	if err != nil {
		return err
	}

	detail.ID = id
	// Set default values for timestamps
	now := time.Now()
	detail.CreatedAt = now
	detail.UpdatedAt = now

	return nil
}

func (r *inquiryPostgresRepository) GetServicePackagePrice(id int) (float64, error) {
	query := `SELECT harga FROM paket_layanan WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var price float64
	err := row.Scan(&price)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("service package with id %d not found", id)
		}
		return 0, err
	}

	return price, nil
}
