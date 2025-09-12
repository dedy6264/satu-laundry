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

	// Start a transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insert into pelanggan table
	customerQuery := `INSERT INTO pelanggan (nama_lengkap, email, nomor_hp, alamat, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_pelanggan`
	fmt.Println("::", customerQuery)
	now := time.Now()
	err = tx.QueryRow(customerQuery, customer.Name, customer.Email, customer.Phone, customer.Address, now, now).Scan(&customer.ID)
	if err != nil {
		return err
	}

	// Insert into pelanggan_outlet junction table
	junctionQuery := `INSERT INTO pelanggan_outlet (id_pelanggan, id_outlet, created_at) VALUES ($1, $2, $3)`
	_, err = tx.Exec(junctionQuery, customer.ID, customer.OutletID, now)
	if err != nil {
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	customer.CreatedAt = now
	customer.UpdatedAt = now
	return nil
}

func (r *customerPostgresRepository) FindByID(id int) (*entities.Customer, error) {
	query := `SELECT p.id_pelanggan, COALESCE(po.id_outlet, 0) as id_outlet, p.nama_lengkap, p.email, p.nomor_hp, p.alamat, p.created_at, p.updated_at 
	          FROM pelanggan p
			  LEFT JOIN pelanggan_outlet po ON p.id_pelanggan = po.id_pelanggan
			  WHERE p.id_pelanggan = $1`

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
	query := `SELECT p.id_pelanggan, po.id_outlet, p.nama_lengkap, p.email, p.nomor_hp, p.alamat, p.created_at, p.updated_at 
	          FROM pelanggan p
			  JOIN pelanggan_outlet po ON p.id_pelanggan = po.id_pelanggan
			  WHERE po.id_outlet = $1`

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
	query := `SELECT p.id_pelanggan, po.id_outlet, p.nama_lengkap, p.email, p.nomor_hp, p.alamat, p.created_at, p.updated_at 
	          FROM pelanggan p
			  JOIN pelanggan_outlet po ON p.id_pelanggan = po.id_pelanggan`

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
	baseQuery := `SELECT p.id_pelanggan, po.id_outlet, p.nama_lengkap, p.email, p.nomor_hp, p.alamat, p.created_at, p.updated_at 
	              FROM pelanggan p
				  JOIN pelanggan_outlet po ON p.id_pelanggan = po.id_pelanggan`
	countQuery := "SELECT COUNT(*) FROM pelanggan p JOIN pelanggan_outlet po ON p.id_pelanggan = po.id_pelanggan"

	var args []interface{}
	argIndex := 1

	// Add search condition if provided
	if search != "" {
		baseQuery += " WHERE LOWER(p.nama_lengkap) LIKE $1 OR LOWER(p.email) LIKE $1 OR p.nomor_hp LIKE $1"
		countQuery += " WHERE LOWER(p.nama_lengkap) LIKE $1 OR LOWER(p.email) LIKE $1 OR p.nomor_hp LIKE $1"
		args = append(args, "%"+search+"%")
		argIndex++
	}

	// Add ordering
	dbOrderBy := "p.id_pelanggan" // default
	switch orderBy {
	case "id_pelanggan":
		dbOrderBy = "p.id_pelanggan"
	case "id_outlet":
		dbOrderBy = "po.id_outlet"
	case "nama_lengkap":
		dbOrderBy = "p.nama_lengkap"
	case "email":
		dbOrderBy = "p.email"
	case "nomor_hp":
		dbOrderBy = "p.nomor_hp"
	case "alamat":
		dbOrderBy = "p.alamat"
	case "created_at":
		dbOrderBy = "p.created_at"
	case "updated_at":
		dbOrderBy = "p.updated_at"
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
	// Start a transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Update pelanggan table
	customerQuery := `UPDATE pelanggan SET nama_lengkap = $1, email = $2, nomor_hp = $3, alamat = $4, updated_at = $5 
	          WHERE id_pelanggan = $6`

	now := time.Now()
	_, err = tx.Exec(customerQuery, customer.Name, customer.Email, customer.Phone, customer.Address, now, customer.ID)
	if err != nil {
		return err
	}

	// Update pelanggan_outlet junction table
	junctionQuery := `UPDATE pelanggan_outlet SET id_outlet = $1 WHERE id_pelanggan = $2`
	_, err = tx.Exec(junctionQuery, customer.OutletID, customer.ID)
	if err != nil {
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	customer.UpdatedAt = now
	return nil
}

func (r *customerPostgresRepository) Delete(id int) error {
	// Start a transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Delete from pelanggan_outlet junction table first (due to foreign key constraints)
	junctionQuery := `DELETE FROM pelanggan_outlet WHERE id_pelanggan = $1`
	_, err = tx.Exec(junctionQuery, id)
	if err != nil {
		return err
	}

	// Delete from pelanggan table
	customerQuery := `DELETE FROM pelanggan WHERE id_pelanggan = $1`
	_, err = tx.Exec(customerQuery, id)
	if err != nil {
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
