package repositories

import (
	"database/sql"
	"fmt"
	"laundry-backend/internal/entities"
)

type transactionPostgresRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionPostgresRepository{
		db: db,
	}
}

func (r *transactionPostgresRepository) FindAll() ([]entities.Transaction, error) {
	query := `
		SELECT 
			t.id_transaksi, t.id_pelanggan, t.id_outlet, t.id_pegawai, t.nomor_invoice, 
			t.tanggal_masuk, t.tanggal_selesai, t.tanggal_diambil, 
			t.total_harga, t.uang_bayar, t.uang_kembalian, t.status_transaksi,
			t.status_pembayaran, t.metode_pembayaran, COALESCE(t.catatan,''), t.status_kode, 
			t.status_pesan, t.nomor_referensi_pembayaran, t.created_at, t.updated_at,
			t.created_by, t.updated_by
		FROM transaksi t
		ORDER BY t.id_transaksi`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entities.Transaction
	for rows.Next() {
		var transaction entities.Transaction
		var employeeID sql.NullInt64
		var entryDate, completionDate, pickupDate sql.NullTime
		var statusCode, statusMessage, paymentReferenceNumber, createdBy, updatedBy sql.NullString

		err := rows.Scan(
			&transaction.ID,
			&transaction.CustomerID,
			&transaction.OutletID,
			&employeeID,
			&transaction.InvoiceNumber,
			&entryDate,
			&completionDate,
			&pickupDate,
			&transaction.TotalPrice,
			&transaction.PaidAmount,
			&transaction.ChangeAmount,
			&transaction.Status,
			&transaction.PaymentStatus,
			&transaction.PaymentMethod,
			&transaction.Note,
			&statusCode,
			&statusMessage,
			&paymentReferenceNumber,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&createdBy,
			&updatedBy,
		)
		if err != nil {
			return nil, err
		}

		// Handle nullable fields
		if employeeID.Valid {
			val := int(employeeID.Int64)
			transaction.EmployeeID = &val
		}
		if entryDate.Valid {
			transaction.EntryDate = &entryDate.Time
		}
		if completionDate.Valid {
			transaction.CompletionDate = &completionDate.Time
		}
		if pickupDate.Valid {
			transaction.PickupDate = &pickupDate.Time
		}
		if statusCode.Valid {
			transaction.StatusCode = &statusCode.String
		}
		if statusMessage.Valid {
			transaction.StatusMessage = &statusMessage.String
		}
		if paymentReferenceNumber.Valid {
			transaction.PaymentReferenceNumber = &paymentReferenceNumber.String
		}
		if createdBy.Valid {
			transaction.CreatedBy = &createdBy.String
		}
		if updatedBy.Valid {
			transaction.UpdatedBy = &updatedBy.String
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *transactionPostgresRepository) FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Transaction, int, error) {
	// Base query
	baseQuery := `
		FROM transaksi t`

	// Count query
	countQuery := "SELECT COUNT(*) " + baseQuery

	// Data query
	dataQuery := `
		SELECT 
			t.id_transaksi, t.id_pelanggan, t.id_outlet, t.id_pegawai, t.nomor_invoice, 
			t.tanggal_masuk, t.tanggal_selesai, t.tanggal_diambil, 
			t.total_harga, t.uang_bayar, t.uang_kembalian, t.status_transaksi,
			t.status_pembayaran, t.metode_pembayaran, COALESCE(t.catatan,''), t.status_kode, 
			t.status_pesan, t.nomor_referensi_pembayaran, t.created_at, t.updated_at,
			t.created_by, t.updated_by
		` + baseQuery

	// Search condition
	var args []interface{}
	if search != "" {
		countQuery += " WHERE t.nomor_invoice ILIKE $1"
		dataQuery += " WHERE t.nomor_invoice ILIKE $1"
		args = append(args, "%"+search+"%")
	}

	// Get total count
	var totalCount int
	countArgs := make([]interface{}, len(args))
	copy(countArgs, args)
	err := r.db.QueryRow(countQuery, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	// Add ordering
	if orderBy == "" {
		orderBy = "t.id_transaksi"
	}
	if orderDir == "" {
		orderDir = "DESC"
	}
	dataQuery += fmt.Sprintf(" ORDER BY %s %s LIMIT $%d OFFSET $%d", orderBy, orderDir, len(args)+1, len(args)+2)

	// Add limit and offset
	args = append(args, limit, offset)

	// Execute data query
	rows, err := r.db.Query(dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var transactions []entities.Transaction
	for rows.Next() {
		var transaction entities.Transaction
		var employeeID sql.NullInt64
		var entryDate, completionDate, pickupDate sql.NullTime
		var statusCode, statusMessage, paymentReferenceNumber, createdBy, updatedBy sql.NullString

		err := rows.Scan(
			&transaction.ID,
			&transaction.CustomerID,
			&transaction.OutletID,
			&employeeID,
			&transaction.InvoiceNumber,
			&entryDate,
			&completionDate,
			&pickupDate,
			&transaction.TotalPrice,
			&transaction.PaidAmount,
			&transaction.ChangeAmount,
			&transaction.Status,
			&transaction.PaymentStatus,
			&transaction.PaymentMethod,
			&transaction.Note,
			&statusCode,
			&statusMessage,
			&paymentReferenceNumber,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&createdBy,
			&updatedBy,
		)
		if err != nil {
			return nil, 0, err
		}

		// Handle nullable fields
		if employeeID.Valid {
			val := int(employeeID.Int64)
			transaction.EmployeeID = &val
		}
		if entryDate.Valid {
			transaction.EntryDate = &entryDate.Time
		}
		if completionDate.Valid {
			transaction.CompletionDate = &completionDate.Time
		}
		if pickupDate.Valid {
			transaction.PickupDate = &pickupDate.Time
		}
		if statusCode.Valid {
			transaction.StatusCode = &statusCode.String
		}
		if statusMessage.Valid {
			transaction.StatusMessage = &statusMessage.String
		}
		if paymentReferenceNumber.Valid {
			transaction.PaymentReferenceNumber = &paymentReferenceNumber.String
		}
		if createdBy.Valid {
			transaction.CreatedBy = &createdBy.String
		}
		if updatedBy.Valid {
			transaction.UpdatedBy = &updatedBy.String
		}

		transactions = append(transactions, transaction)
	}

	return transactions, totalCount, nil
}

func (r *transactionPostgresRepository) FindByID(id int) (*entities.Transaction, error) {
	query := `
		SELECT 
			t.id_transaksi, t.id_pelanggan, t.id_outlet, t.id_pegawai, t.nomor_invoice, 
			t.tanggal_masuk, t.tanggal_selesai, t.tanggal_diambil, 
			t.total_harga, t.uang_bayar, t.uang_kembalian, t.status_transaksi,
			t.status_pembayaran, t.metode_pembayaran, COALESCE(t.catatan,''), t.status_kode, 
			t.status_pesan, t.nomor_referensi_pembayaran, t.created_at, t.updated_at,
			t.created_by, t.updated_by
		FROM transaksi t
		WHERE t.id_transaksi = $1`

	var transaction entities.Transaction
	var employeeID sql.NullInt64
	var entryDate, completionDate, pickupDate sql.NullTime
	var statusCode, statusMessage, paymentReferenceNumber, createdBy, updatedBy sql.NullString

	err := r.db.QueryRow(query, id).Scan(
		&transaction.ID,
		&transaction.CustomerID,
		&transaction.OutletID,
		&employeeID,
		&transaction.InvoiceNumber,
		&entryDate,
		&completionDate,
		&pickupDate,
		&transaction.TotalPrice,
		&transaction.PaidAmount,
		&transaction.ChangeAmount,
		&transaction.Status,
		&transaction.PaymentStatus,
		&transaction.PaymentMethod,
		&transaction.Note,
		&statusCode,
		&statusMessage,
		&paymentReferenceNumber,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
		&createdBy,
		&updatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Handle nullable fields
	if employeeID.Valid {
		val := int(employeeID.Int64)
		transaction.EmployeeID = &val
	}
	if entryDate.Valid {
		transaction.EntryDate = &entryDate.Time
	}
	if completionDate.Valid {
		transaction.CompletionDate = &completionDate.Time
	}
	if pickupDate.Valid {
		transaction.PickupDate = &pickupDate.Time
	}
	if statusCode.Valid {
		transaction.StatusCode = &statusCode.String
	}
	if statusMessage.Valid {
		transaction.StatusMessage = &statusMessage.String
	}
	if paymentReferenceNumber.Valid {
		transaction.PaymentReferenceNumber = &paymentReferenceNumber.String
	}
	if createdBy.Valid {
		transaction.CreatedBy = &createdBy.String
	}
	if updatedBy.Valid {
		transaction.UpdatedBy = &updatedBy.String
	}

	return &transaction, nil
}

func (r *transactionPostgresRepository) FindByOutletID(outletID int) ([]entities.Transaction, error) {
	query := `
		SELECT 
			t.id_transaksi, t.id_pelanggan, t.id_outlet, t.id_pegawai, t.nomor_invoice, 
			t.tanggal_masuk, t.tanggal_selesai, t.tanggal_diambil, 
			t.total_harga, t.uang_bayar, t.uang_kembalian, t.status_transaksi,
			t.status_pembayaran, t.metode_pembayaran, COALESCE(t.catatan,''), t.status_kode, 
			t.status_pesan, t.nomor_referensi_pembayaran, t.created_at, t.updated_at,
			t.created_by, t.updated_by
		FROM transaksi t
		WHERE t.id_outlet = $1
		ORDER BY t.id_transaksi`

	rows, err := r.db.Query(query, outletID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entities.Transaction
	for rows.Next() {
		var transaction entities.Transaction
		var employeeID sql.NullInt64
		var entryDate, completionDate, pickupDate sql.NullTime
		var statusCode, statusMessage, paymentReferenceNumber, createdBy, updatedBy sql.NullString

		err := rows.Scan(
			&transaction.ID,
			&transaction.CustomerID,
			&transaction.OutletID,
			&employeeID,
			&transaction.InvoiceNumber,
			&entryDate,
			&completionDate,
			&pickupDate,
			&transaction.TotalPrice,
			&transaction.PaidAmount,
			&transaction.ChangeAmount,
			&transaction.Status,
			&transaction.PaymentStatus,
			&transaction.PaymentMethod,
			&transaction.Note,
			&statusCode,
			&statusMessage,
			&paymentReferenceNumber,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&createdBy,
			&updatedBy,
		)
		if err != nil {
			return nil, err
		}

		// Handle nullable fields
		if employeeID.Valid {
			val := int(employeeID.Int64)
			transaction.EmployeeID = &val
		}
		if entryDate.Valid {
			transaction.EntryDate = &entryDate.Time
		}
		if completionDate.Valid {
			transaction.CompletionDate = &completionDate.Time
		}
		if pickupDate.Valid {
			transaction.PickupDate = &pickupDate.Time
		}
		if statusCode.Valid {
			transaction.StatusCode = &statusCode.String
		}
		if statusMessage.Valid {
			transaction.StatusMessage = &statusMessage.String
		}
		if paymentReferenceNumber.Valid {
			transaction.PaymentReferenceNumber = &paymentReferenceNumber.String
		}
		if createdBy.Valid {
			transaction.CreatedBy = &createdBy.String
		}
		if updatedBy.Valid {
			transaction.UpdatedBy = &updatedBy.String
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *transactionPostgresRepository) FindDetailsByTransactionID(transactionID int) ([]entities.TransactionDetail, error) {
	query := `
		SELECT 
			td.id_detail, td.id_transaksi, td.id_layanan, td.kuantitas, td.harga_satuan, td.subtotal,
			td.created_at, td.updated_at, td.created_by, td.updated_by
		FROM detail_transaksi td
		WHERE td.id_transaksi = $1
		ORDER BY td.id_detail`

	rows, err := r.db.Query(query, transactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []entities.TransactionDetail
	for rows.Next() {
		var detail entities.TransactionDetail
		var quantity, price, subtotal sql.NullFloat64
		var createdBy, updatedBy sql.NullString

		err := rows.Scan(
			&detail.ID,
			&detail.TransactionID,
			&detail.ServiceID,
			&quantity,
			&price,
			&subtotal,
			&detail.CreatedAt,
			&detail.UpdatedAt,
			&createdBy,
			&updatedBy,
		)
		if err != nil {
			return nil, err
		}

		// Handle nullable fields
		if quantity.Valid {
			detail.Quantity = &quantity.Float64
		}
		if price.Valid {
			detail.Price = &price.Float64
		}
		if subtotal.Valid {
			detail.Subtotal = &subtotal.Float64
		}
		if createdBy.Valid {
			detail.CreatedBy = &createdBy.String
		}
		if updatedBy.Valid {
			detail.UpdatedBy = &updatedBy.String
		}

		details = append(details, detail)
	}

	return details, nil
}
