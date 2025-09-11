package repositories

import (
	"database/sql"
	"laundry-backend/internal/entities"
	"fmt"
	"strings"
)

type outletPostgresRepository struct {
	db *sql.DB
}

func NewOutletRepository(db *sql.DB) OutletRepository {
	return &outletPostgresRepository{db: db}
}

func (r *outletPostgresRepository) Create(outlet *entities.Outlet) error {
	// Handle nullable float fields
	var lat, lon interface{}
	if outlet.Latitude != nil {
		lat = *outlet.Latitude
	}
	if outlet.Longitude != nil {
		lon = *outlet.Longitude
	}

	query := `INSERT INTO outlet (id_cabang, nama_outlet, alamat, kota, provinsi, kode_pos, telepon, email, 
		latitude, longitude, jam_buka, jam_tutup, pic_nama, pic_email, pic_telepon, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, NOW(), NOW()) RETURNING id_outlet`
	return r.db.QueryRow(query, outlet.CabangID, outlet.Name, outlet.Address, outlet.City, outlet.Province, 
		outlet.PostalCode, outlet.Phone, outlet.Email, lat, lon, outlet.OpenTime, 
		outlet.CloseTime, outlet.PICName, outlet.PICEmail, outlet.PICTelepon).Scan(&outlet.ID)
}

func (r *outletPostgresRepository) FindByID(id int) (*entities.Outlet, error) {
	query := `SELECT id_outlet, id_cabang, nama_outlet, alamat, kota, provinsi, kode_pos, telepon, email, 
		latitude, longitude, jam_buka, jam_tutup, pic_nama, pic_email, pic_telepon, created_at, updated_at 
	FROM outlet WHERE id_outlet = $1`
	row := r.db.QueryRow(query, id)

	var outlet entities.Outlet
	var lat, lon sql.NullFloat64
	err := row.Scan(
		&outlet.ID,
		&outlet.CabangID,
		&outlet.Name,
		&outlet.Address,
		&outlet.City,
		&outlet.Province,
		&outlet.PostalCode,
		&outlet.Phone,
		&outlet.Email,
		&lat,
		&lon,
		&outlet.OpenTime,
		&outlet.CloseTime,
		&outlet.PICName,
		&outlet.PICEmail,
		&outlet.PICTelepon,
		&outlet.CreatedAt,
		&outlet.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Handle nullable float fields
	if lat.Valid {
		outlet.Latitude = &lat.Float64
	}
	if lon.Valid {
		outlet.Longitude = &lon.Float64
	}

	return &outlet, nil
}

func (r *outletPostgresRepository) FindByCabangID(cabangID int) ([]entities.Outlet, error) {
	query := `SELECT id_outlet, id_cabang, nama_outlet, alamat, kota, provinsi, kode_pos, telepon, email, 
		latitude, longitude, jam_buka, jam_tutup, pic_nama, pic_email, pic_telepon, created_at, updated_at 
	FROM outlet WHERE id_cabang = $1`
	rows, err := r.db.Query(query, cabangID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var outlets []entities.Outlet
	for rows.Next() {
		var outlet entities.Outlet
		var lat, lon sql.NullFloat64
		err := rows.Scan(
			&outlet.ID,
			&outlet.CabangID,
			&outlet.Name,
			&outlet.Address,
			&outlet.City,
			&outlet.Province,
			&outlet.PostalCode,
			&outlet.Phone,
			&outlet.Email,
			&lat,
			&lon,
			&outlet.OpenTime,
			&outlet.CloseTime,
			&outlet.PICName,
			&outlet.PICEmail,
			&outlet.PICTelepon,
			&outlet.CreatedAt,
			&outlet.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Handle nullable float fields
		if lat.Valid {
			outlet.Latitude = &lat.Float64
		}
		if lon.Valid {
			outlet.Longitude = &lon.Float64
		}

		outlets = append(outlets, outlet)
	}

	return outlets, nil
}

func (r *outletPostgresRepository) FindAll() ([]entities.Outlet, error) {
	query := `SELECT id_outlet, id_cabang, nama_outlet, alamat, kota, provinsi, kode_pos, telepon, email, 
		latitude, longitude, jam_buka, jam_tutup, pic_nama, pic_email, pic_telepon, created_at, updated_at 
	FROM outlet`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var outlets []entities.Outlet
	for rows.Next() {
		var outlet entities.Outlet
		var lat, lon sql.NullFloat64
		err := rows.Scan(
			&outlet.ID,
			&outlet.CabangID,
			&outlet.Name,
			&outlet.Address,
			&outlet.City,
			&outlet.Province,
			&outlet.PostalCode,
			&outlet.Phone,
			&outlet.Email,
			&lat,
			&lon,
			&outlet.OpenTime,
			&outlet.CloseTime,
			&outlet.PICName,
			&outlet.PICEmail,
			&outlet.PICTelepon,
			&outlet.CreatedAt,
			&outlet.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Handle nullable float fields
		if lat.Valid {
			outlet.Latitude = &lat.Float64
		}
		if lon.Valid {
			outlet.Longitude = &lon.Float64
		}

		outlets = append(outlets, outlet)
	}

	return outlets, nil
}

func (r *outletPostgresRepository) FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Outlet, int, int, error) {
	// Validate orderBy field to prevent SQL injection
	validFields := map[string]bool{
		"id":          true,
		"cabang_id":   true,
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
		"id":          "id_outlet",
		"cabang_id":   "id_cabang",
		"name":        "nama_outlet",
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
	baseQuery := `SELECT id_outlet, id_cabang, nama_outlet, alamat, kota, provinsi, kode_pos, telepon, email, 
		latitude, longitude, jam_buka, jam_tutup, pic_nama, pic_email, pic_telepon, created_at, updated_at FROM outlet`
	countQuery := `SELECT COUNT(*) FROM outlet`
	
	var args []interface{}
	argIndex := 1
	
	// Add search condition if provided
	if search != "" {
		search = strings.ToLower(search)
		baseQuery += fmt.Sprintf(` WHERE LOWER(nama_outlet) LIKE $%d OR LOWER(kota) LIKE $%d OR LOWER(provinsi) LIKE $%d OR LOWER(pic_nama) LIKE $%d`, argIndex, argIndex+1, argIndex+2, argIndex+3)
		countQuery += fmt.Sprintf(` WHERE LOWER(nama_outlet) LIKE $%d OR LOWER(kota) LIKE $%d OR LOWER(provinsi) LIKE $%d OR LOWER(pic_nama) LIKE $%d`, argIndex, argIndex+1, argIndex+2, argIndex+3)
		args = append(args, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
		argIndex += 4
	}
	
	// Add ordering
	dbOrderBy := fieldMap[orderBy]
	if dbOrderBy == "" {
		dbOrderBy = "id_outlet"
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

	var outlets []entities.Outlet
	for rows.Next() {
		var outlet entities.Outlet
		var lat, lon sql.NullFloat64
		err := rows.Scan(
			&outlet.ID,
			&outlet.CabangID,
			&outlet.Name,
			&outlet.Address,
			&outlet.City,
			&outlet.Province,
			&outlet.PostalCode,
			&outlet.Phone,
			&outlet.Email,
			&lat,
			&lon,
			&outlet.OpenTime,
			&outlet.CloseTime,
			&outlet.PICName,
			&outlet.PICEmail,
			&outlet.PICTelepon,
			&outlet.CreatedAt,
			&outlet.UpdatedAt,
		)
		if err != nil {
			return nil, 0, 0, err
		}

		// Handle nullable float fields
		if lat.Valid {
			outlet.Latitude = &lat.Float64
		}
		if lon.Valid {
			outlet.Longitude = &lon.Float64
		}

		outlets = append(outlets, outlet)
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
	
	return outlets, recordsTotal, recordsFiltered, nil
}

func (r *outletPostgresRepository) Update(outlet *entities.Outlet) error {
	// Handle nullable float fields
	var lat, lon interface{}
	if outlet.Latitude != nil {
		lat = *outlet.Latitude
	}
	if outlet.Longitude != nil {
		lon = *outlet.Longitude
	}

	query := `UPDATE outlet SET id_cabang = $1, nama_outlet = $2, alamat = $3, kota = $4, provinsi = $5, 
		kode_pos = $6, telepon = $7, email = $8, latitude = $9, longitude = $10, jam_buka = $11, jam_tutup = $12, 
		pic_nama = $13, pic_email = $14, pic_telepon = $15, updated_at = NOW() WHERE id_outlet = $16`
	_, err := r.db.Exec(query, outlet.CabangID, outlet.Name, outlet.Address, outlet.City, outlet.Province, 
		outlet.PostalCode, outlet.Phone, outlet.Email, lat, lon, outlet.OpenTime, 
		outlet.CloseTime, outlet.PICName, outlet.PICEmail, outlet.PICTelepon, outlet.ID)
	return err
}

func (r *outletPostgresRepository) Delete(id int) error {
	query := `DELETE FROM outlet WHERE id_outlet = $1`
	_, err := r.db.Exec(query, id)
	return err
}