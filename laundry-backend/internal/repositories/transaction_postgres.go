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
		t.id_transaksi,
		t.id_pelanggan,
		t.id_outlet,
		t.id_access,
		t.nomor_invoice,			 
		t.tanggal_masuk,
		t.tanggal_selesai,
		t.tanggal_diambil,			 
		t.total_harga,
		t.uang_bayar,
		t.uang_kembalian,
		t.status_transaksi,			
		COALESCE(t.catatan,''),
		t.created_at,
		t.updated_at,			
		t.created_by,
		t.updated_by
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
		var userID sql.NullInt64
		var entryDate, completionDate, pickupDate sql.NullTime
		var createdBy, updatedBy sql.NullString

		err := rows.Scan(
			&transaction.ID,
			&transaction.CustomerID,
			&transaction.OutletID,
			&userID,
			&transaction.InvoiceNumber,
			&entryDate,
			&completionDate,
			&pickupDate,
			&transaction.TotalPrice,
			&transaction.PaidAmount,
			&transaction.ChangeAmount,
			&transaction.Status,
			&transaction.Note,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&createdBy,
			&updatedBy,
		)
		if err != nil {
			return nil, err
		}

		// Handle nullable fields
		if userID.Valid {
			val := int(userID.Int64)
			transaction.UserID = &val
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
		t.id_transaksi,
		t.id_pelanggan,
		t.id_outlet,
		t.id_access,
		t.nomor_invoice,			 
		t.tanggal_masuk,
		t.tanggal_selesai,
		t.tanggal_diambil,			 
		t.total_harga,
		t.uang_bayar,
		t.uang_kembalian,
		t.status_transaksi,			
		COALESCE(t.catatan,''),
		t.created_at,
		t.updated_at,			
		t.created_by,
		t.updated_by
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
		var userID sql.NullInt64
		var entryDate, completionDate, pickupDate sql.NullTime
		var createdBy, updatedBy sql.NullString

		err := rows.Scan(
			&transaction.ID,
			&transaction.CustomerID,
			&transaction.OutletID,
			&userID,
			&transaction.InvoiceNumber,
			&entryDate,
			&completionDate,
			&pickupDate,
			&transaction.TotalPrice,
			&transaction.PaidAmount,
			&transaction.ChangeAmount,
			&transaction.Status,
			&transaction.Note,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&createdBy,
			&updatedBy,
		)
		if err != nil {
			return nil, 0, err
		}

		// Handle nullable fields
		if userID.Valid {
			val := int(userID.Int64)
			transaction.UserID = &val
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
			t.id_transaksi,
		t.id_pelanggan,
		t.id_outlet,
		t.id_access,
		t.nomor_invoice,			 
		t.tanggal_masuk,
		t.tanggal_selesai,
		t.tanggal_diambil,			 
		t.total_harga,
		t.uang_bayar,
		t.uang_kembalian,
		t.status_transaksi,			
		COALESCE(t.catatan,''),
		t.created_at,
		t.updated_at,			
		t.created_by,
		t.updated_by
		FROM transaksi t
		WHERE t.id_transaksi = $1`

	var transaction entities.Transaction
	var userID sql.NullInt64
	var entryDate, completionDate, pickupDate sql.NullTime
	var createdBy, updatedBy sql.NullString

	err := r.db.QueryRow(query, id).Scan(
		&transaction.ID,
		&transaction.CustomerID,
		&transaction.OutletID,
		&userID,
		&transaction.InvoiceNumber,
		&entryDate,
		&completionDate,
		&pickupDate,
		&transaction.TotalPrice,
		&transaction.PaidAmount,
		&transaction.ChangeAmount,
		&transaction.Status,
		&transaction.Note,
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
	if userID.Valid {
		val := int(userID.Int64)
		transaction.UserID = &val
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
			t.id_transaksi,
		t.id_pelanggan,
		t.id_outlet,
		t.id_access,
		t.nomor_invoice,			 
		t.tanggal_masuk,
		t.tanggal_selesai,
		t.tanggal_diambil,			 
		t.total_harga,
		t.uang_bayar,
		t.uang_kembalian,
		t.status_transaksi,			
		COALESCE(t.catatan,''),
		t.created_at,
		t.updated_at,			
		t.created_by,
		t.updated_by
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
		var userID sql.NullInt64
		var entryDate, completionDate, pickupDate sql.NullTime
		var createdBy, updatedBy sql.NullString

		err := rows.Scan(
			&transaction.ID,
			&transaction.CustomerID,
			&transaction.OutletID,
			&userID,
			&transaction.InvoiceNumber,
			&entryDate,
			&completionDate,
			&pickupDate,
			&transaction.TotalPrice,
			&transaction.PaidAmount,
			&transaction.ChangeAmount,
			&transaction.Status,
			&transaction.Note,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&createdBy,
			&updatedBy,
		)
		if err != nil {
			return nil, err
		}

		// Handle nullable fields
		if userID.Valid {
			val := int(userID.Int64)
			transaction.UserID = &val
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
			td.id_detail,
			td.id_transaksi,
			td.id_layanan,
			td.kuantitas,
			td.harga_satuan,
			td.subtotal,
			td.status_pengerjaan,
			td.created_at,
			td.updated_at,
			td.created_by,
			td.updated_by
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
			&detail.Status,
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

func (r *transactionPostgresRepository) UpdateTransactionStatus(id int, status string) error {
	query := `
		UPDATE transaksi
		SET status_transaksi = $1, updated_at = NOW()
		WHERE id_transaksi = $2`

	_, err := r.db.Exec(query, status, id)
	return err
}

// ////perlu update
func (r *transactionPostgresRepository) UpdatePaymentStatus(id int, status string) error {
	query := `
		UPDATE transaksi
		SET status_pembayaran = $1, updated_at = NOW()
		WHERE id_transaksi = $2`

	_, err := r.db.Exec(query, status, id)
	return err
}

func (r *transactionPostgresRepository) UpdatePaymentCallback(transactionID int, request entities.PaymentCallbackRequest) error {
	query := `
		UPDATE transaksi
		SET 
			status_pembayaran = $1,
			metode_pembayaran = $2,
			nomor_referensi_pembayaran = $3,
			uang_bayar = $4,
			uang_kembalian = $5,
			status_kode = $6,
			status_pesan = $7,
			updated_at = NOW()
		WHERE id_transaksi = $8`

	_, err := r.db.Exec(query,
		request.PaymentStatus,
		request.PaymentMethod,
		request.PaymentReferenceNumber,
		request.PaidAmount,
		request.ChangeAmount,
		request.StatusCode,
		request.StatusMessage,
		transactionID,
	)
	return err
}
