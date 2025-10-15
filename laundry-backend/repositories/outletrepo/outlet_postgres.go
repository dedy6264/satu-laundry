package outletrepo

import (
	"database/sql"
	"laundry-backend/entities"
	"laundry-backend/repositories"
	"strconv"
)

type outletPostgresRepository struct {
	repo repositories.Repositories
}

func NewOutletRepo(repo repositories.Repositories) outletPostgresRepository {
	return outletPostgresRepository{repo: repo}
}

func (r outletPostgresRepository) Create(outlet entities.Outlet) error {
	// Handle nullable float fields
	var lat, lon interface{}
	if outlet.Latitude != nil {
		lat = outlet.Latitude
	}
	if outlet.Longitude != nil {
		lon = outlet.Longitude
	}

	query := `INSERT INTO outlet (id_cabang, nama_outlet, alamat, kota, provinsi, kode_pos, telepon, email, 
		latitude, longitude, jam_buka, jam_tutup, pic_nama, pic_email, pic_telepon, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, NOW(), NOW()) RETURNING id_outlet`
	return r.repo.Db.QueryRow(query, outlet.CabangID, outlet.Name, outlet.Address, outlet.City, outlet.Province,
		outlet.PostalCode, outlet.Phone, outlet.Email, lat, lon, outlet.OpenTime,
		outlet.CloseTime, outlet.PICName, outlet.PICEmail, outlet.PICTelepon).Scan(&outlet.ID)
}

func (r outletPostgresRepository) FindByID(id int) (entities.Outlet, error) {
	query := `SELECT id_outlet, id_cabang, nama_outlet, alamat, kota, provinsi, kode_pos, telepon, email, 
		latitude, longitude, jam_buka, jam_tutup, pic_nama, pic_email, pic_telepon, created_at, updated_at 
	FROM outlet WHERE id_outlet = $1`
	row := r.repo.Db.QueryRow(query, id)

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
			return outlet, nil
		}
		return outlet, err
	}

	// Handle nullable float fields
	if lat.Valid {
		outlet.Latitude = &lat.Float64
	}
	if lon.Valid {
		outlet.Longitude = &lon.Float64
	}

	return outlet, nil
}

func (r outletPostgresRepository) FindByCabangID(cabangID int) ([]entities.Outlet, error) {
	query := `SELECT id_outlet, id_cabang, nama_outlet, alamat, kota, provinsi, kode_pos, telepon, email, 
		latitude, longitude, jam_buka, jam_tutup, pic_nama, pic_email, pic_telepon, created_at, updated_at 
	FROM outlet WHERE id_cabang = $1`
	rows, err := r.repo.Db.Query(query, cabangID)
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

func (r outletPostgresRepository) FindAll(request entities.Outlet) ([]entities.Outlet, error) {
	query := `SELECT 
	a.id_outlet,
	a.id_cabang,
	b.nama_cabang,
	c.id_brand,
	c.nama_brand,
	a.nama_outlet,
	a.alamat,
	a.kota,
	a.provinsi,
	a.kode_pos,
	a.telepon,
	a.email,	 
	a.latitude,
	a.longitude,
	a.jam_buka,
	a.jam_tutup,
	a.pic_nama,
	a.pic_email,
	a.pic_telepon,
	a.created_at,
	a.updated_at 
	FROM outlet as a
	leftjoin cabang as b on b.id = a.id_cabang
	leftjoin brand as c on c.id = b.id_brand where true `
	if request.CabangID != 0 {
		query += ` and id_cabang = ` + strconv.Itoa(request.CabangID)
	}
	if request.ID != 0 {
		query += ` and id_outlet = ` + strconv.Itoa(request.ID)
	}

	rows, err := r.repo.Db.Query(query)
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
			&outlet.NamaCabang,
			&outlet.BrandID,
			&outlet.NamaBrand,
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

func (r outletPostgresRepository) FindAllWithPagination(request entities.DTRequest[entities.Outlet]) ([]entities.Outlet, int, error) {

	// Build the query
	baseQuery := `SELECT id_outlet, id_cabang, nama_outlet, alamat, kota, provinsi, kode_pos, telepon, email, 
		latitude, longitude, jam_buka, jam_tutup, pic_nama, pic_email, pic_telepon, created_at, updated_at FROM outlet where true `
	countQuery := `SELECT COUNT() FROM outlet`
	if request.Data.ID != 0 {
		baseQuery += ` and id_outlet = ` + strconv.Itoa(request.Data.ID)
	}
	if request.Data.CabangID != 0 {
		baseQuery += ` and id_cabang = ` + strconv.Itoa(request.Data.CabangID)
	}
	if request.OrderBy != "" {
		baseQuery += ` ORDER BY ` + request.OrderBy + ` ` + request.SortBy
	} else {
		baseQuery += ` ORDER BY id_outlet ASC`
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
			return nil, 0, err
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
	var recordsTotal int
	err = r.repo.Db.QueryRow(countQuery).Scan(&recordsTotal)
	if err != nil {
		return nil, 0, err
	}

	return outlets, recordsTotal, nil
}

func (r outletPostgresRepository) Update(outlet entities.Outlet) error {
	// Handle nullable float fields
	var lat, lon interface{}
	if outlet.Latitude != nil {
		lat = outlet.Latitude
	}
	if outlet.Longitude != nil {
		lon = outlet.Longitude
	}

	query := `UPDATE outlet SET id_cabang = $1, nama_outlet = $2, alamat = $3, kota = $4, provinsi = $5, 
		kode_pos = $6, telepon = $7, email = $8, latitude = $9, longitude = $10, jam_buka = $11, jam_tutup = $12, 
		pic_nama = $13, pic_email = $14, pic_telepon = $15, updated_at = NOW() WHERE id_outlet = $16`
	_, err := r.repo.Db.Exec(query, outlet.CabangID, outlet.Name, outlet.Address, outlet.City, outlet.Province,
		outlet.PostalCode, outlet.Phone, outlet.Email, lat, lon, outlet.OpenTime,
		outlet.CloseTime, outlet.PICName, outlet.PICEmail, outlet.PICTelepon, outlet.ID)
	return err
}

func (r outletPostgresRepository) Delete(id int) error {
	query := `DELETE FROM outlet WHERE id_outlet = $1`
	_, err := r.repo.Db.Exec(query, id)
	return err
}
