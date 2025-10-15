package repositories

import (
	"database/sql"
	"fmt"
	"laundry-backend/internal/entities"
	"strconv"
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
	query := `SELECT 
	a.id_pegawai,
	a.id_outlet,
	b.id_cabang,
	c.id_brand,
	a.nik, 
	a.nama_lengkap, 
	a.email, 
	a.telepon, 
	a.alamat, 
	a.tanggal_lahir, 
	a.jenis_kelamin, 
	a.posisi, 
	a.gaji, 
	a.tanggal_masuk, 
	a.status, 
	a.created_at, 
	a.updated_at FROM pegawai as a 
	join outlet as b on b.id_outlet=a.id_outlet
	join cabang as c on c.id_cabang=b.id_cabang
	join brand as d on d.id_brand=c.id_brand
	WHERE id_pegawai = $1`
	row := r.db.QueryRow(query, id)

	var employee entities.Employee
	err := row.Scan(
		&employee.ID,
		&employee.OutletID,
		&employee.CabangID,
		&employee.BrandID,
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
		fmt.Println(query)

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

func (r *employeePostgresRepository) FindAllWithPagination(request entities.DTRequest[entities.Employee]) ([]entities.Employee, int, error) {

	// Build the query
	baseQuery := `SELECT 
	a.id_pegawai,
	a.id_outlet,
	b.id_cabang,
	c.id_brand,
	a.nik,
	a.nama_lengkap,
	a.email,
	a.telepon,
	a.alamat,
	a.tanggal_lahir,
	a.jenis_kelamin,
	a.posisi,
	a.gaji,
	a.tanggal_masuk,
	a.status,
	a.created_at,
	a.updated_at FROM pegawai as a
	join outlet as b on b.id_outlet=a.id_outlet
	join cabang as c on c.id_cabang=b.id_cabang
	join brand as d on d.id_brand=c.id_brand where true`
	countQuery := `SELECT COUNT(*) FROM pegawai`
	if request.Data.ID != 0 {
		baseQuery += ` and a.id_pegawai = ` + strconv.Itoa(request.Data.ID)
	}
	if request.Data.OutletID != 0 {
		baseQuery += ` and a.id_outlet = ` + strconv.Itoa(request.Data.OutletID)
	}
	if request.Data.CabangID != 0 {
		baseQuery += ` and c.id_cabang = ` + strconv.Itoa(request.Data.CabangID)
	}
	if request.Data.BrandID != 0 {
		baseQuery += ` and d.id_brand = ` + strconv.Itoa(request.Data.BrandID)
	}
	if request.OrderBy != "" {
		baseQuery += ` ORDER BY ` + request.OrderBy + ` ` + request.SortBy
	} else {
		baseQuery += ` ORDER BY a.id_pegawai ASC`
	}
	if request.Length != 0 {
		baseQuery += ` LIMIT ` + strconv.Itoa(request.Length) + ` OFFSET ` + strconv.Itoa(request.Start)
	}
	// Execute the data query
	rows, err := r.db.Query(baseQuery)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var employees []entities.Employee
	for rows.Next() {
		var employee entities.Employee
		err := rows.Scan(
			&employee.ID,
			&employee.OutletID,
			&employee.CabangID,
			&employee.BrandID,
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
			return nil, 0, err
		}

		employees = append(employees, employee)
	}

	// Execute the count query
	var recordsTotal int
	err = r.db.QueryRow(countQuery).Scan(&recordsTotal)
	if err != nil {
		return nil, 0, err
	}

	return employees, recordsTotal, nil
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
