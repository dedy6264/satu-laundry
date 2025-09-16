package repositories

import (
	"database/sql"
	"fmt"
	"laundry-backend/internal/entities"
	"laundry-backend/internal/utils"
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

func (r *inquiryPostgresRepository) ValidateEmployee(id int) (*entities.Employee, error) {
	// Use the employee repository to find the employee by ID
	employee, err := r.employeeRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// If employee is not found, return nil
	if employee == nil {
		return nil, nil
	}

	// Employee exists, return the employee data
	return employee, nil
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

func (r *inquiryPostgresRepository) GetServicePackagePrice(id int) (float64, error) {
	query := `SELECT harga_satuan FROM paket_layanan WHERE id_layanan = $1`
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

// Transaction methods
func (r *inquiryPostgresRepository) BeginTransaction() (*sql.Tx, error) {
	return r.db.Begin()
}

func (r *inquiryPostgresRepository) InsertTransactionWithTx(tx *sql.Tx, transaction *entities.Transaction) (int, error) {
	query := `INSERT INTO transaksi (

		id_pelanggan, 
		id_outlet, 
		nomor_invoice, 
		tanggal_masuk, 
		status_transaksi, 
		catatan,
		id_pegawai,
		tanggal_selesai,
		tanggal_diambil,
		total_harga,
		uang_bayar,
		uang_kembalian,
		created_at,
		updated_at,
		created_by,
		updated_by
	) VALUES (?, ?, ?, ?, ?, ?,?,?,?,?,?,?,?,?,?,?
	) RETURNING id_transaksi`

	var id int
	query = utils.QuerySupport(query)
	err := tx.QueryRow(
		query,
		transaction.CustomerID,
		transaction.OutletID,
		transaction.InvoiceNumber,
		transaction.EntryDate,
		transaction.Status,
		transaction.Note,
		transaction.EmployeeID,
		transaction.CompletionDate,
		transaction.PickupDate,
		transaction.TotalPrice,
		transaction.PaidAmount,
		transaction.ChangeAmount,
		transaction.CreatedAt,
		transaction.UpdatedAt,
		transaction.CreatedBy,
		transaction.UpdatedBy,
	).Scan(&id)

	if err != nil {
		return id, err
	}

	transaction.ID = id
	// Set default values for timestamps
	now := time.Now()
	transaction.CreatedAt = now
	transaction.UpdatedAt = now

	return id, nil
}

func (r *inquiryPostgresRepository) InsertTransactionDetailWithTx(tx *sql.Tx, detail *entities.TransactionDetail) error {
	query := `INSERT INTO detail_transaksi (
			id_transaksi,
			id_layanan,
			kuantitas,
			harga_satuan,
			subtotal,
			status_pengerjaan,
			created_at,
			updated_at,
			created_by,
			updated_by
	) VALUES (?,?,?,?,?,?,?,?,?,?
	) RETURNING id_detail`

	var id int
	query = utils.QuerySupport(query)
	err := tx.QueryRow(
		query,
		detail.TransactionID,
		detail.ServiceID,
		detail.Quantity,
		detail.Price,
		detail.Subtotal,
		detail.Status,
		detail.CreatedAt,
		detail.UpdatedAt,
		detail.CreatedBy,
		detail.UpdatedBy,
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

func (r *inquiryPostgresRepository) InsertPaymentWithTx(tx *sql.Tx, payment *entities.Payment) error {
	query := `INSERT INTO pembayaran (
			id_transaksi,
			tanggal_bayar,
			jumlah_bayar,
			metode_bayar,
			nomor_referensi_partner,
			status_code_partner,
			status_message_partner,
			catatan,
			created_at,
			updated_at
	) VALUES (?,?,?,?,?,?,?,?,?,?
	) RETURNING id_pembayaran`

	var id int
	query = utils.QuerySupport(query)
	err := tx.QueryRow(
		query,
		payment.TransactionID,
		payment.PaymentDate,
		payment.Amount,
		payment.Method,
		payment.PartnerReferenceNo,
		payment.PartnerStatusCode,
		payment.PartnerStatusMessage,
		payment.Note,
		payment.CreatedAt,
		payment.UpdatedAt,
	).Scan(&id)

	if err != nil {
		return err
	}

	payment.ID = id
	// Set default values for timestamps
	now := time.Now()
	payment.CreatedAt = now
	payment.UpdatedAt = now

	return nil
}

func (r *inquiryPostgresRepository) InsertHistoryStatusTransactionWithTx(tx *sql.Tx, history *entities.HistoryStatusTransaction) error {
	query := `INSERT INTO history_status_transaksi (
			id_transaksi,
			status_lama,
			status_baru,
			waktu_perubahan,
			keterangan,
			created_at,
			updated_at
	) VALUES (?,?,?,?,?,?,?
	) RETURNING id_history`

	var id int
	query = utils.QuerySupport(query)
	err := tx.QueryRow(
		query,
		history.TransactionID,
		history.OldStatus,
		history.NewStatus,
		history.ChangeTime,
		history.Description,
		history.CreatedAt,
		history.UpdatedAt,
	).Scan(&id)

	if err != nil {
		return err
	}

	history.ID = id
	// Set default values for timestamps
	now := time.Now()
	history.CreatedAt = now
	history.UpdatedAt = now

	return nil
}
