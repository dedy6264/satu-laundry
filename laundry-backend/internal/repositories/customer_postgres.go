package repositories

import (
	"database/sql"
	"laundry-backend/internal/entities"

	"fmt"
	"time"
)

type customerPostgresRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerPostgresRepository{db: db}
}

func (r *customerPostgresRepository) Create(customer *entities.Customer) error {
	query := `INSERT INTO pelanggan (id_outlet, nama, email, telepon, alamat, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id_pelanggan`

	now := time.Now()
	err := r.db.QueryRow(query, customer.OutletID, customer.Name, customer.Email, customer.Phone, customer.Address, now, now).Scan(&customer.ID)
	if err != nil {
		return err
	}

	customer.CreatedAt = now
	customer.UpdatedAt = now
	return nil
}

func (r *customerPostgresRepository) FindByID(id int) (*entities.Customer, error) {
	query := `SELECT id_pelanggan, id_outlet, nama, email, telepon, alamat, created_at, updated_at 
	          FROM pelanggan WHERE id_pelanggan = $1`

	row := r.db.QueryRow(query, id)

	customer := &entities.Customer{}
	err := row.Scan(&customer.ID, &customer.OutletID, &customer.Name, &customer.Email,
		&customer.Phone, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return customer, nil
}

func (r *customerPostgresRepository) FindByOutletID(outletID int) ([]entities.Customer, error) {
	query := `SELECT id_pelanggan, id_outlet, nama, email, telepon, alamat, created_at, updated_at 
	          FROM pelanggan WHERE id_outlet = $1`

	rows, err := r.db.Query(query, outletID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []entities.Customer
	for rows.Next() {
		customer := &entities.Customer{}
		err := rows.Scan(&customer.ID, &customer.OutletID, &customer.Name, &customer.Email,
			&customer.Phone, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return nil, err
		}
		customers = append(customers, *customer)
	}

	return customers, nil
}

func (r *customerPostgresRepository) FindAll() ([]entities.Customer, error) {
	query := `SELECT id_pelanggan, id_outlet, nama, email, telepon, alamat, created_at, updated_at 
	          FROM pelanggan`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []entities.Customer
	for rows.Next() {
		customer := &entities.Customer{}
		err := rows.Scan(&customer.ID, &customer.OutletID, &customer.Name, &customer.Email,
			&customer.Phone, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return nil, err
		}
		customers = append(customers, *customer)
	}

	return customers, nil
}

func (r *customerPostgresRepository) FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Customer, int, int, error) {
	baseQuery := `SELECT id_pelanggan, id_outlet, nama, email, telepon, alamat, created_at, updated_at 
	              FROM pelanggan`
	countQuery := "SELECT COUNT(*) FROM pelanggan"

	var args []interface{}
	argIndex := 1

	// Add search condition if provided
	if search != "" {
		baseQuery += " WHERE LOWER(nama) LIKE $1 OR LOWER(email) LIKE $1 OR telepon LIKE $1"
		countQuery += " WHERE LOWER(nama) LIKE $1 OR LOWER(email) LIKE $1 OR telepon LIKE $1"
		args = append(args, "%"+search+"%")
		argIndex++
	}

	// Add ordering
	dbOrderBy := "id_pelanggan" // default
	switch orderBy {
	case "id_pelanggan":
		dbOrderBy = "id_pelanggan"
	case "id_outlet":
		dbOrderBy = "id_outlet"
	case "nama":
		dbOrderBy = "nama"
	case "email":
		dbOrderBy = "email"
	case "telepon":
		dbOrderBy = "telepon"
	case "alamat":
		dbOrderBy = "alamat"
	case "created_at":
		dbOrderBy = "created_at"
	case "updated_at":
		dbOrderBy = "updated_at"
	}

	// Validate order direction
	if orderDir != "asc" && orderDir != "desc" {
		orderDir = "asc"
	}

	baseQuery += " ORDER BY " + dbOrderBy + " " + orderDir

	// Add pagination
	baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, limit, offset)

	// Execute the data query
	rows, err := r.db.Query(baseQuery, args...)
	if err != nil {
		return nil, 0, 0, err
	}
	defer rows.Close()

	var customers []entities.Customer
	for rows.Next() {
		customer := &entities.Customer{}
		err := rows.Scan(&customer.ID, &customer.OutletID, &customer.Name, &customer.Email,
			&customer.Phone, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return nil, 0, 0, err
		}
		customers = append(customers, *customer)
	}

	// Execute the count query
	var recordsTotal, recordsFiltered int
	err = r.db.QueryRow(countQuery, args[:len(args)-2]...).Scan(&recordsTotal)
	if err != nil {
		return nil, 0, 0, err
	}

	// If search is applied, we need to get the filtered count
	if search != "" && len(args[:len(args)-2]) > 0 {
		searchArgs := args[:len(args)-2] // Remove limit and offset args
		err = r.db.QueryRow(countQuery, searchArgs...).Scan(&recordsFiltered)
		if err != nil {
			return nil, 0, 0, err
		}
	} else {
		recordsFiltered = recordsTotal
	}

	return customers, recordsTotal, recordsFiltered, nil
}

func (r *customerPostgresRepository) Update(customer *entities.Customer) error {
	query := `UPDATE pelanggan SET id_outlet = $1, nama = $2, email = $3, telepon = $4, alamat = $5, updated_at = $6 
	          WHERE id_pelanggan = $7`

	now := time.Now()
	_, err := r.db.Exec(query, customer.OutletID, customer.Name, customer.Email, customer.Phone, customer.Address, now, customer.ID)
	if err != nil {
		return err
	}

	customer.UpdatedAt = now
	return nil
}

func (r *customerPostgresRepository) Delete(id int) error {
	query := `DELETE FROM pelanggan WHERE id_pelanggan = $1`
	_, err := r.db.Exec(query, id)
	return err
}
