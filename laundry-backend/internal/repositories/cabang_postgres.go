package repositories

import (
	"database/sql"
	"laundry-backend/internal/entities"
	"fmt"
	"strings"
)

type cabangPostgresRepository struct {
	db *sql.DB
}

func NewCabangRepository(db *sql.DB) CabangRepository {
	return &cabangPostgresRepository{db: db}
}

func (r *cabangPostgresRepository) Create(cabang *entities.Cabang) error {
	query := `INSERT INTO cabang (id_brand, nama_cabang, alamat, kota, provinsi, kode_pos, telepon, email, 
		pic_nama, pic_email, pic_telepon, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW(), NOW()) RETURNING id_cabang`
	return r.db.QueryRow(query, cabang.BrandID, cabang.Name, cabang.Address, cabang.City, cabang.Province, 
		cabang.PostalCode, cabang.Phone, cabang.Email, cabang.PICName, cabang.PICEmail, cabang.PICTelepon).Scan(&cabang.ID)
}

func (r *cabangPostgresRepository) FindByID(id int) (*entities.Cabang, error) {
	query := `SELECT id_cabang, id_brand, nama_cabang, alamat, kota, provinsi, kode_pos, telepon, email, 
		pic_nama, pic_email, pic_telepon, created_at, updated_at 
	FROM cabang WHERE id_cabang = $1`
	row := r.db.QueryRow(query, id)

	var cabang entities.Cabang
	err := row.Scan(
		&cabang.ID,
		&cabang.BrandID,
		&cabang.Name,
		&cabang.Address,
		&cabang.City,
		&cabang.Province,
		&cabang.PostalCode,
		&cabang.Phone,
		&cabang.Email,
		&cabang.PICName,
		&cabang.PICEmail,
		&cabang.PICTelepon,
		&cabang.CreatedAt,
		&cabang.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &cabang, nil
}

func (r *cabangPostgresRepository) FindByBrandID(brandID int) ([]entities.Cabang, error) {
	query := `SELECT id_cabang, id_brand, nama_cabang, alamat, kota, provinsi, kode_pos, telepon, email, 
		pic_nama, pic_email, pic_telepon, created_at, updated_at 
	FROM cabang WHERE id_brand = $1`
	rows, err := r.db.Query(query, brandID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cabangs []entities.Cabang
	for rows.Next() {
		var cabang entities.Cabang
		err := rows.Scan(
			&cabang.ID,
			&cabang.BrandID,
			&cabang.Name,
			&cabang.Address,
			&cabang.City,
			&cabang.Province,
			&cabang.PostalCode,
			&cabang.Phone,
			&cabang.Email,
			&cabang.PICName,
			&cabang.PICEmail,
			&cabang.PICTelepon,
			&cabang.CreatedAt,
			&cabang.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		cabangs = append(cabangs, cabang)
	}

	return cabangs, nil
}

func (r *cabangPostgresRepository) FindAll() ([]entities.Cabang, error) {
	query := `SELECT id_cabang, id_brand, nama_cabang, alamat, kota, provinsi, kode_pos, telepon, email, 
		pic_nama, pic_email, pic_telepon, created_at, updated_at 
	FROM cabang`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cabangs []entities.Cabang
	for rows.Next() {
		var cabang entities.Cabang
		err := rows.Scan(
			&cabang.ID,
			&cabang.BrandID,
			&cabang.Name,
			&cabang.Address,
			&cabang.City,
			&cabang.Province,
			&cabang.PostalCode,
			&cabang.Phone,
			&cabang.Email,
			&cabang.PICName,
			&cabang.PICEmail,
			&cabang.PICTelepon,
			&cabang.CreatedAt,
			&cabang.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		cabangs = append(cabangs, cabang)
	}

	return cabangs, nil
}

func (r *cabangPostgresRepository) FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Cabang, int, int, error) {
	// Validate orderBy field to prevent SQL injection
	validFields := map[string]bool{
		"id":          true,
		"brand_id":    true,
		"name":        true,
		"city":        true,
		"province":    true,
		"pic_name":    true,
		"pic_email":   true,
		"created_at":  true,
		"updated_at":  true,
	}
	
	// Map field names to database column names
	fieldMap := map[string]string{
		"id":          "id_cabang",
		"brand_id":    "id_brand",
		"name":        "nama_cabang",
		"city":        "kota",
		"province":    "provinsi",
		"pic_name":    "pic_nama",
		"pic_email":   "pic_email",
		"created_at":  "created_at",
		"updated_at":  "updated_at",
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
	baseQuery := `SELECT id_cabang, id_brand, nama_cabang, alamat, kota, provinsi, kode_pos, telepon, email, 
		pic_nama, pic_email, pic_telepon, created_at, updated_at FROM cabang`
	countQuery := `SELECT COUNT(*) FROM cabang`
	
	var args []interface{}
	argIndex := 1
	
	// Add search condition if provided
	if search != "" {
		search = strings.ToLower(search)
		baseQuery += fmt.Sprintf(` WHERE LOWER(nama_cabang) LIKE $%d OR LOWER(kota) LIKE $%d OR LOWER(provinsi) LIKE $%d OR LOWER(pic_nama) LIKE $%d`, argIndex, argIndex+1, argIndex+2, argIndex+3)
		countQuery += fmt.Sprintf(` WHERE LOWER(nama_cabang) LIKE $%d OR LOWER(kota) LIKE $%d OR LOWER(provinsi) LIKE $%d OR LOWER(pic_nama) LIKE $%d`, argIndex, argIndex+1, argIndex+2, argIndex+3)
		args = append(args, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
		argIndex += 4
	}
	
	// Add ordering
	dbOrderBy := fieldMap[orderBy]
	if dbOrderBy == "" {
		dbOrderBy = "id_cabang"
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

	var cabangs []entities.Cabang
	for rows.Next() {
		var cabang entities.Cabang
		err := rows.Scan(
			&cabang.ID,
			&cabang.BrandID,
			&cabang.Name,
			&cabang.Address,
			&cabang.City,
			&cabang.Province,
			&cabang.PostalCode,
			&cabang.Phone,
			&cabang.Email,
			&cabang.PICName,
			&cabang.PICEmail,
			&cabang.PICTelepon,
			&cabang.CreatedAt,
			&cabang.UpdatedAt,
		)
		if err != nil {
			return nil, 0, 0, err
		}
		cabangs = append(cabangs, cabang)
	}
	
	// Execute the count query
	var recordsTotal, recordsFiltered int
	err = r.db.QueryRow(countQuery).Scan(&recordsTotal)
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
	
	return cabangs, recordsTotal, recordsFiltered, nil
}

func (r *cabangPostgresRepository) Update(cabang *entities.Cabang) error {
	query := `UPDATE cabang SET id_brand = $1, nama_cabang = $2, alamat = $3, kota = $4, provinsi = $5, 
		kode_pos = $6, telepon = $7, email = $8, pic_nama = $9, pic_email = $10, pic_telepon = $11, 
		updated_at = NOW() WHERE id_cabang = $12`
	_, err := r.db.Exec(query, cabang.BrandID, cabang.Name, cabang.Address, cabang.City, cabang.Province, 
		cabang.PostalCode, cabang.Phone, cabang.Email, cabang.PICName, cabang.PICEmail, cabang.PICTelepon, cabang.ID)
	return err
}

func (r *cabangPostgresRepository) Delete(id int) error {
	query := `DELETE FROM cabang WHERE id_cabang = $1`
	_, err := r.db.Exec(query, id)
	return err
}