package repositories

import (
	"database/sql"
	"laundry-backend/internal/entities"
	"strconv"
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

func (r *cabangPostgresRepository) FindAllWithPagination(request entities.DTRequest[entities.Cabang]) ([]entities.Cabang, int, error) {

	// Build the query
	baseQuery := `SELECT 
	id_cabang, 
	id_brand, 
	nama_cabang, 
	alamat, 
	kota, 
	provinsi, 
	kode_pos, 
	telepon, 
	email, 
	pic_nama, 
	pic_email, 
	pic_telepon, 
	created_at, 
	updated_at FROM cabang where true`
	countQuery := `SELECT COUNT(*) FROM cabang`
	if request.Data.ID != 0 {
		baseQuery += ` and id_cabang = ` + strconv.Itoa(request.Data.ID)
	}
	if request.Data.BrandID != 0 {
		baseQuery += ` and id_brand = ` + strconv.Itoa(request.Data.BrandID)
	}
	if request.OrderBy != "" {
		baseQuery += ` ORDER BY ` + request.OrderBy + ` ` + request.SortBy
	} else {
		baseQuery += ` ORDER BY id_cabang ASC`
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
			return nil, 0, err
		}
		cabangs = append(cabangs, cabang)
	}

	// Execute the count query
	var recordsTotal int
	err = r.db.QueryRow(countQuery).Scan(&recordsTotal)
	if err != nil {
		return nil, 0, err
	}

	return cabangs, recordsTotal, nil
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
