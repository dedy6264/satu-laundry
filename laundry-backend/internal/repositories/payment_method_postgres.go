package repositories

import (
	"database/sql"
	"fmt"
	"laundry-backend/internal/entities"
	"strings"
)

type paymentMethodPostgresRepository struct {
	db *sql.DB
}

func NewPaymentMethodRepository(db *sql.DB) PaymentMethodRepository {
	return &paymentMethodPostgresRepository{db: db}
}

func (r *paymentMethodPostgresRepository) Create(paymentMethod *entities.PaymentMethod) error {
	query := `INSERT INTO metode_pembayaran (nama_metode, url, s_key, m_key, merchant_fee, admin_fee, status, created_at, updated_at,  created_by, updated_by) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW(),  $8, $9) RETURNING id`
	return r.db.QueryRow(query, paymentMethod.NamaMetode, paymentMethod.URL, paymentMethod.SKey, paymentMethod.MKey, paymentMethod.MerchantFee, paymentMethod.AdminFee, paymentMethod.Status, paymentMethod.CreatedBy, paymentMethod.UpdatedBy).Scan(&paymentMethod.ID)
}

func (r *paymentMethodPostgresRepository) FindByID(id int) (*entities.PaymentMethod, error) {
	query := `SELECT id, nama_metode, url, s_key, m_key, merchant_fee, admin_fee, status, created_at, updated_at,  created_by, updated_by 
	FROM metode_pembayaran WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var paymentMethod entities.PaymentMethod
	err := row.Scan(
		&paymentMethod.ID,
		&paymentMethod.NamaMetode,
		&paymentMethod.URL,
		&paymentMethod.SKey,
		&paymentMethod.MKey,
		&paymentMethod.MerchantFee,
		&paymentMethod.AdminFee,
		&paymentMethod.Status,
		&paymentMethod.CreatedAt,
		&paymentMethod.UpdatedAt,
		&paymentMethod.CreatedBy,
		&paymentMethod.UpdatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &paymentMethod, nil
}

func (r *paymentMethodPostgresRepository) FindAll() ([]entities.PaymentMethod, error) {
	query := `SELECT id, nama_metode, url, s_key, m_key, merchant_fee, admin_fee, status, created_at, updated_at,  created_by, updated_by FROM metode_pembayaran`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paymentMethods []entities.PaymentMethod
	for rows.Next() {
		var paymentMethod entities.PaymentMethod
		err := rows.Scan(
			&paymentMethod.ID,
			&paymentMethod.NamaMetode,
			&paymentMethod.URL,
			&paymentMethod.SKey,
			&paymentMethod.MKey,
			&paymentMethod.MerchantFee,
			&paymentMethod.AdminFee,
			&paymentMethod.Status,
			&paymentMethod.CreatedAt,
			&paymentMethod.UpdatedAt,
			&paymentMethod.CreatedBy,
			&paymentMethod.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}
		paymentMethods = append(paymentMethods, paymentMethod)
	}

	return paymentMethods, nil
}

func (r *paymentMethodPostgresRepository) FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.PaymentMethod, int, int, error) {
	// Validate orderBy field to prevent SQL injection
	validFields := map[string]bool{
		"id":           true,
		"nama_metode":  true,
		"url":          true,
		"merchant_fee": true,
		"admin_fee":    true,
		"status":       true,
		"created_at":   true,
		"updated_at":   true,
	}

	// Default to id if orderBy is not valid
	if !validFields[orderBy] {
		orderBy = "id"
	}

	// Default to asc if orderDir is not valid
	if orderDir != "asc" && orderDir != "desc" {
		orderDir = "asc"
	}

	// Build the query
	baseQuery := `SELECT id, nama_metode, url, s_key, m_key, merchant_fee, admin_fee, status, created_at, updated_at,  created_by, updated_by FROM metode_pembayaran`
	countQuery := `SELECT COUNT(*) FROM metode_pembayaran`

	var args []interface{}
	argIndex := 1

	// Add search condition if provided
	if search != "" {
		search = strings.ToLower(search)
		baseQuery += fmt.Sprintf(` WHERE LOWER(nama_metode) LIKE $%d`, argIndex)
		countQuery += fmt.Sprintf(` WHERE LOWER(nama_metode) LIKE $%d`, argIndex)
		args = append(args, "%"+search+"%")
		argIndex++
	}

	// Add ordering
	baseQuery += fmt.Sprintf(` ORDER BY %s %s`, orderBy, strings.ToUpper(orderDir))

	// Add pagination
	baseQuery += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, argIndex, argIndex+1)
	args = append(args, limit, offset)

	// Execute the data query
	rows, err := r.db.Query(baseQuery, args...)
	if err != nil {
		return nil, 0, 0, err
	}
	defer rows.Close()

	var paymentMethods []entities.PaymentMethod
	for rows.Next() {
		var paymentMethod entities.PaymentMethod
		err := rows.Scan(
			&paymentMethod.ID,
			&paymentMethod.NamaMetode,
			&paymentMethod.URL,
			&paymentMethod.SKey,
			&paymentMethod.MKey,
			&paymentMethod.MerchantFee,
			&paymentMethod.AdminFee,
			&paymentMethod.Status,
			&paymentMethod.CreatedAt,
			&paymentMethod.UpdatedAt,
			&paymentMethod.CreatedBy,
			&paymentMethod.UpdatedBy,
		)
		if err != nil {
			return nil, 0, 0, err
		}
		paymentMethods = append(paymentMethods, paymentMethod)
	}

	// Execute the count query
	var recordsTotal, recordsFiltered int
	err = r.db.QueryRow(countQuery, args[:len(args)-2]...).Scan(&recordsTotal)
	if err != nil {
		return nil, 0, 0, err
	}

	// If search is applied, we need to get the filtered count
	if search != "" {
		searchArgs := args[:len(args)-2] // Remove limit and offset args
		err = r.db.QueryRow(countQuery, searchArgs...).Scan(&recordsFiltered)
		if err != nil {
			return nil, 0, 0, err
		}
	} else {
		recordsFiltered = recordsTotal
	}

	return paymentMethods, recordsTotal, recordsFiltered, nil
}

func (r *paymentMethodPostgresRepository) Update(paymentMethod *entities.PaymentMethod) error {
	query := `UPDATE metode_pembayaran SET nama_metode = $1, url = $2, s_key = $3, m_key = $4, merchant_fee = $5, admin_fee = $6, 
	status = $7, updated_at = NOW(), created_by = $8, updated_by = $9 WHERE id = $10`
	_, err := r.db.Exec(query, paymentMethod.NamaMetode, paymentMethod.URL, paymentMethod.SKey, paymentMethod.MKey, paymentMethod.MerchantFee, paymentMethod.AdminFee, paymentMethod.Status, paymentMethod.CreatedBy, paymentMethod.UpdatedBy, paymentMethod.ID)
	return err
}

func (r *paymentMethodPostgresRepository) Delete(id int) error {
	// Instead of deleting, we'll set the deleted_at timestamp
	query := `UPDATE metode_pembayaran SET deleted_at = NOW(), status = 'inactive' WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
