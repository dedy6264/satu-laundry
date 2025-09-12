package repositories

import (
	"database/sql"
	"fmt"
	"laundry-backend/internal/entities"
	"strings"
)

type employeePostgresRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeePostgresRepository{db: db}
}

func (r *employeePostgresRepository) Create(employee *entities.Employee) error {
	query := `INSERT INTO pegawai (id_outlet, nik, nama_lengkap, email, telepon, alamat, tanggal_lahir, jenis_kelamin, posisi, gaji, tanggal_masuk, status,  created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, NOW(), NOW()) RETURNING id_pegawai`

	return r.db.QueryRow(query, employee.OutletID, employee.NIK, employee.Name, employee.Email, employee.Phone, employee.Address, employee.BirthDate, employee.Gender, employee.Position, employee.Salary, employee.JoinDate, employee.Status).Scan(&employee.ID)
}

func (r *employeePostgresRepository) FindByID(id int) (*entities.Employee, error) {
	query := `SELECT id_pegawai, id_outlet, nik, nama_lengkap, email, telepon, alamat, tanggal_lahir, jenis_kelamin, posisi, gaji, tanggal_masuk, status,  created_at, updated_at FROM pegawai WHERE id_pegawai = $1`
	row := r.db.QueryRow(query, id)

	var employee entities.Employee
	err := row.Scan(
		&employee.ID,
		&employee.OutletID,
		&employee.NIK,
		&employee.Name,
		&employee.Email,
		&employee.Phone,
		&employee.Address,
		&employee.BirthDate,
		&employee.Gender,
		&employee.Position,
		&employee.Salary,
		&employee.JoinDate,
		&employee.Status,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &employee, nil
}

func (r *employeePostgresRepository) FindAll() ([]entities.Employee, error) {
	query := `SELECT id_pegawai, id_outlet, nik, nama_lengkap, email, telepon, alamat, tanggal_lahir, jenis_kelamin, posisi, gaji, tanggal_masuk, status,  created_at, updated_at FROM pegawai`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []entities.Employee
	for rows.Next() {
		var employee entities.Employee
		err := rows.Scan(
			&employee.ID,
			&employee.OutletID,
			&employee.NIK,
			&employee.Name,
			&employee.Email,
			&employee.Phone,
			&employee.Address,
			&employee.BirthDate,
			&employee.Gender,
			&employee.Position,
			&employee.Salary,
			&employee.JoinDate,
			&employee.Status,
			&employee.CreatedAt,
			&employee.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

func (r *employeePostgresRepository) FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Employee, int, int, error) {
	// Validate orderBy field to prevent SQL injection
	validFields := map[string]bool{
		"id_pegawai":    true,
		"id_outlet":     true,
		"nik":           true,
		"nama_lengkap":  true,
		"email":         true,
		"telepon":       true,
		"alamat":        true,
		"tanggal_lahir": true,
		"jenis_kelamin": true,
		"posisi":        true,
		"gaji":          true,
		"tanggal_masuk": true,
		"status":        true,
		"created_at":    true,
		"updated_at":    true,
	}

	// Map field names to database column names
	fieldMap := map[string]string{
		"id_pegawai":    "id_pegawai",
		"id_outlet":     "id_outlet",
		"nik":           "nik",
		"nama_lengkap":  "nama_lengkap",
		"email":         "email",
		"telepon":       "telepon",
		"alamat":        "alamat",
		"tanggal_lahir": "tanggal_lahir",
		"jenis_kelamin": "jenis_kelamin",
		"posisi":        "posisi",
		"gaji":          "gaji",
		"tanggal_masuk": "tanggal_masuk",
		"status":        "status",
		"created_at":    "created_at",
		"updated_at":    "updated_at",
	}

	// Default to id_pegawai if orderBy is not valid
	if !validFields[orderBy] {
		orderBy = "id_pegawai"
	}

	// Default to asc if orderDir is not valid
	if orderDir != "asc" && orderDir != "desc" {
		orderDir = "asc"
	}

	// Build the query
	baseQuery := `SELECT id_pegawai, id_outlet, nik, nama_lengkap, email, telepon, alamat, tanggal_lahir, jenis_kelamin, posisi, gaji, tanggal_masuk, status, created_at, updated_at FROM pegawai`
	countQuery := `SELECT COUNT(*) FROM pegawai`

	var args []interface{}
	argIndex := 1

	// Add search condition if provided
	if search != "" {
		search = strings.ToLower(search)
		baseQuery += fmt.Sprintf(` WHERE (LOWER(nik) LIKE $%d OR LOWER(nama_lengkap) LIKE $%d OR LOWER(email) LIKE $%d)`, argIndex, argIndex+1, argIndex+2)
		countQuery += fmt.Sprintf(` WHERE (LOWER(nik) LIKE $%d OR LOWER(nama_lengkap) LIKE $%d OR LOWER(email) LIKE $%d)`, argIndex, argIndex+1, argIndex+2)
		args = append(args, "%"+search+"%", "%"+search+"%", "%"+search+"%")
		argIndex += 3
	}

	// Add ordering
	dbOrderBy := fieldMap[orderBy]
	if dbOrderBy == "" {
		dbOrderBy = "id_pegawai"
	}
	baseQuery += fmt.Sprintf(` ORDER BY %s %s`, dbOrderBy, strings.ToUpper(orderDir))

	// Add pagination
	baseQuery += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, argIndex, argIndex+1)
	args = append(args, limit, offset)

	// Execute the data query
	rows, err := r.db.Query(baseQuery, args...)
	if err != nil {
		return nil, 0, 0, err
	}
	defer rows.Close()

	var employees []entities.Employee
	for rows.Next() {
		var employee entities.Employee
		err := rows.Scan(
			&employee.ID,
			&employee.OutletID,
			&employee.NIK,
			&employee.Name,
			&employee.Email,
			&employee.Phone,
			&employee.Address,
			&employee.BirthDate,
			&employee.Gender,
			&employee.Position,
			&employee.Salary,
			&employee.JoinDate,
			&employee.Status,
			&employee.CreatedAt,
			&employee.UpdatedAt,
		)
		if err != nil {
			return nil, 0, 0, err
		}

		employees = append(employees, employee)
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

	return employees, recordsTotal, recordsFiltered, nil
}

func (r *employeePostgresRepository) Update(employee *entities.Employee) error {
	query := `UPDATE pegawai SET id_outlet = $1, nik = $2, nama_lengkap = $3, email = $4, telepon = $5, alamat = $6, tanggal_lahir = $7, jenis_kelamin = $8, posisi = $9, gaji = $10, tanggal_masuk = $11, status = $12,  updated_at = NOW() WHERE id_pegawai = $13`
	_, err := r.db.Exec(query, employee.OutletID, employee.NIK, employee.Name, employee.Email, employee.Phone, employee.Address, employee.BirthDate, employee.Gender, employee.Position, employee.Salary, employee.JoinDate, employee.Status, employee.ID)
	return err
}

func (r *employeePostgresRepository) Delete(id int) error {
	query := `DELETE FROM pegawai WHERE id_pegawai = $1`
	_, err := r.db.Exec(query, id)
	return err
}
