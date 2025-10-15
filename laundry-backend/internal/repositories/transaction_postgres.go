package repositories

import (
	"database/sql"
	"laundry-backend/internal/entities"
	"strconv"
)

type transactionPostgresRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionPostgresRepository{
		db: db,
	}
}

func (r *transactionPostgresRepository) FindAll() (response []entities.Transaction, err error) {
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
			transaction.UserID = val
		}
		if entryDate.Valid {
			transaction.EntryDate = entryDate.Time
		}
		if completionDate.Valid {
			transaction.CompletionDate = completionDate.Time
		}
		if pickupDate.Valid {
			transaction.PickupDate = pickupDate.Time
		}

		if createdBy.Valid {
			transaction.CreatedBy = createdBy.String
		}
		if updatedBy.Valid {
			transaction.UpdatedBy = updatedBy.String
		}

		response = append(response, transaction)
	}

	return response, nil
}

func (r *transactionPostgresRepository) FindAllWithPagination(request entities.DTRequest[entities.Transaction]) (response []entities.Transaction, totalCount int, err error) {
	// Base query
	baseQuery := `
		FROM transaksi where true `

	// Count query
	countQuery := "SELECT COUNT(*) " + baseQuery

	// Data query
	getQuery := `
		SELECT 
		id_transaksi,
		id_pelanggan,
		id_outlet,
		id_access,
		nomor_invoice,			 
		tanggal_masuk,
		tanggal_selesai,
		tanggal_diambil,			 
		total_harga,
		uang_bayar,
		uang_kembalian,
		status_transaksi,			
		COALESCE(catatan,''),
		created_at,
		updated_at,			
		created_by,
		updated_by
		` + baseQuery

	if request.Data.ID != 0 {
		getQuery += ` and id_transaksi = ` + strconv.Itoa(request.Data.ID)
	}
	if request.Data.CustomerID != 0 {
		getQuery += ` and id_pelanggan = ` + strconv.Itoa(request.Data.CustomerID)
	}
	if request.OrderBy != "" {
		getQuery += ` ORDER BY ` + request.OrderBy + ` ` + request.SortBy
	} else {
		getQuery += ` ORDER BY id_transaksi ASC`
	}
	if request.Length != 0 {
		getQuery += ` LIMIT ` + strconv.Itoa(request.Length) + ` OFFSET ` + strconv.Itoa(request.Start)
	}
	// Get total count
	err = r.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}
	// Execute data query
	rows, err := r.db.Query(getQuery)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

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
			transaction.UserID = val
		}
		if entryDate.Valid {
			transaction.EntryDate = entryDate.Time
		}
		if completionDate.Valid {
			transaction.CompletionDate = completionDate.Time
		}
		if pickupDate.Valid {
			transaction.PickupDate = pickupDate.Time
		}
		if createdBy.Valid {
			transaction.CreatedBy = createdBy.String
		}
		if updatedBy.Valid {
			transaction.UpdatedBy = updatedBy.String
		}

		response = append(response, transaction)
	}

	return response, totalCount, nil
}

func (r *transactionPostgresRepository) FindByID(id int) (response entities.Transaction, err error) {
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

	err = r.db.QueryRow(query, id).Scan(
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
			return response, nil
		}
		return response, err
	}

	// Handle nullable fields
	if userID.Valid {
		val := int(userID.Int64)
		transaction.UserID = val
	}
	if entryDate.Valid {
		transaction.EntryDate = entryDate.Time
	}
	if completionDate.Valid {
		transaction.CompletionDate = completionDate.Time
	}
	if pickupDate.Valid {
		transaction.PickupDate = pickupDate.Time
	}

	if createdBy.Valid {
		transaction.CreatedBy = createdBy.String
	}
	if updatedBy.Valid {
		transaction.UpdatedBy = updatedBy.String
	}

	return response, nil
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
			transaction.UserID = val
		}
		if entryDate.Valid {
			transaction.EntryDate = entryDate.Time
		}
		if completionDate.Valid {
			transaction.CompletionDate = completionDate.Time
		}
		if pickupDate.Valid {
			transaction.PickupDate = pickupDate.Time
		}
		if createdBy.Valid {
			transaction.CreatedBy = createdBy.String
		}
		if updatedBy.Valid {
			transaction.UpdatedBy = updatedBy.String
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
			detail.Quantity = quantity.Float64
		}
		if price.Valid {
			detail.Price = price.Float64
		}
		if subtotal.Valid {
			detail.Subtotal = subtotal.Float64
		}
		if createdBy.Valid {
			detail.CreatedBy = createdBy.String
		}
		if updatedBy.Valid {
			detail.UpdatedBy = updatedBy.String
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
