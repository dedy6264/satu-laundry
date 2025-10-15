package paymentmethodrepo

import (
	"database/sql"
	"laundry-backend/entities"
	"laundry-backend/repositories"
	"strconv"
)

type paymentMethodPostgresRepository struct {
	repo repositories.Repositories
}

func NewPaymentMethodRepo(repo repositories.Repositories) paymentMethodPostgresRepository {
	return paymentMethodPostgresRepository{repo: repo}
}

func (r paymentMethodPostgresRepository) Create(paymentMethod entities.PaymentMethod) error {
	query := `INSERT INTO metode_pembayaran (nama_metode, url, s_key, m_key, merchant_fee, admin_fee, status, created_at, updated_at,  created_by, updated_by) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW(),  $8, $9) RETURNING id`
	return r.repo.Db.QueryRow(query, paymentMethod.NamaMetode, paymentMethod.URL, paymentMethod.SKey, paymentMethod.MKey, paymentMethod.MerchantFee, paymentMethod.AdminFee, paymentMethod.Status, paymentMethod.CreatedBy, paymentMethod.UpdatedBy).Scan(&paymentMethod.ID)
}

func (r paymentMethodPostgresRepository) FindByID(id int) (entities.PaymentMethod, error) {
	query := `SELECT id, nama_metode, url, s_key, m_key, merchant_fee, admin_fee, status, created_at, updated_at,  created_by, updated_by 
	FROM metode_pembayaran WHERE id = $1`
	row := r.repo.Db.QueryRow(query, id)

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
			return paymentMethod, nil
		}
		return paymentMethod, err
	}

	return paymentMethod, nil
}

func (r paymentMethodPostgresRepository) FindAll() ([]entities.PaymentMethod, error) {
	query := `SELECT id, nama_metode, url, s_key, m_key, merchant_fee, admin_fee, status, created_at, updated_at,  created_by, updated_by FROM metode_pembayaran`
	rows, err := r.repo.Db.Query(query)
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

func (r paymentMethodPostgresRepository) FindAllWithPagination(request entities.DTRequest[entities.PaymentMethod]) ([]entities.PaymentMethod, int, error) {

	// Build the query
	baseQuery := `SELECT id, nama_metode, url, s_key, m_key, merchant_fee, admin_fee, status, created_at, updated_at,  created_by, updated_by FROM metode_pembayaran where true `
	countQuery := `SELECT COUNT(*) FROM metode_pembayaran`
	if request.Data.ID != 0 {
		baseQuery += ` and id = ` + strconv.Itoa(request.Data.ID)
	}
	if request.OrderBy != "" {
		baseQuery += ` ORDER BY ` + request.OrderBy + ` ` + request.SortBy
	} else {
		baseQuery += ` ORDER BY id ASC`
	}
	if request.Length != 0 {
		baseQuery += ` LIMIT ` + strconv.Itoa(request.Length) + ` OFFSET ` + strconv.Itoa(request.Start)
	}
	// Execute the data query
	rows, err := r.repo.Db.Query(baseQuery)
	if err != nil {
		return nil, 0, err
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
			return nil, 0, err
		}
		paymentMethods = append(paymentMethods, paymentMethod)
	}

	// Execute the count query
	var recordsTotal int
	err = r.repo.Db.QueryRow(countQuery).Scan(&recordsTotal)
	if err != nil {
		return nil, 0, err
	}

	return paymentMethods, recordsTotal, nil
}

func (r paymentMethodPostgresRepository) Update(paymentMethod entities.PaymentMethod) error {
	query := `UPDATE metode_pembayaran SET nama_metode = $1, url = $2, s_key = $3, m_key = $4, merchant_fee = $5, admin_fee = $6, 
	status = $7, updated_at = NOW(), created_by = $8, updated_by = $9 WHERE id = $10`
	_, err := r.repo.Db.Exec(query, paymentMethod.NamaMetode, paymentMethod.URL, paymentMethod.SKey, paymentMethod.MKey, paymentMethod.MerchantFee, paymentMethod.AdminFee, paymentMethod.Status, paymentMethod.CreatedBy, paymentMethod.UpdatedBy, paymentMethod.ID)
	return err
}

func (r paymentMethodPostgresRepository) Delete(id int) error {
	// Instead of deleting, we'll set the deleted_at timestamp
	query := `UPDATE metode_pembayaran SET deleted_at = NOW(), status = 'inactive' WHERE id = $1`
	_, err := r.repo.Db.Exec(query, id)
	return err
}
