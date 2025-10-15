package cabangrepo

import (
	"database/sql"
	"laundry-backend/entities"
	"laundry-backend/repositories"
	"strconv"
)

type cabangPostgresRepository struct {
	repo repositories.Repositories
}

func NewCabangRepo(repo repositories.Repositories) cabangPostgresRepository {
	return cabangPostgresRepository{repo: repo}
}

func (r cabangPostgresRepository) Create(cabang entities.Cabang) error {
	query := `INSERT INTO cabang (id_brand, nama_cabang, alamat, kota, provinsi, kode_pos, telepon, email, 
		pic_nama, pic_email, pic_telepon, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW(), NOW()) RETURNING id_cabang`
	return r.repo.Db.QueryRow(query, cabang.BrandID, cabang.Name, cabang.Address, cabang.City, cabang.Province,
		cabang.PostalCode, cabang.Phone, cabang.Email, cabang.PICName, cabang.PICEmail, cabang.PICTelepon).Scan(&cabang.ID)
}

func (r cabangPostgresRepository) FindByID(id int) (response entities.Cabang, err error) {
	query := `SELECT id_cabang, id_brand, nama_cabang, alamat, kota, provinsi, kode_pos, telepon, email, 
		pic_nama, pic_email, pic_telepon, created_at, updated_at 
	FROM cabang WHERE id_cabang = $1`
	row := r.repo.Db.QueryRow(query, id)

	err = row.Scan(
		&response.ID,
		&response.BrandID,
		&response.Name,
		&response.Address,
		&response.City,
		&response.Province,
		&response.PostalCode,
		&response.Phone,
		&response.Email,
		&response.PICName,
		&response.PICEmail,
		&response.PICTelepon,
		&response.CreatedAt,
		&response.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return response, nil
		}
		return response, err
	}

	return response, nil
}

func (r cabangPostgresRepository) FindByBrandID(brandID int) (response []entities.Cabang, err error) {
	query := `SELECT id_cabang, id_brand, nama_cabang, alamat, kota, provinsi, kode_pos, telepon, email, 
		pic_nama, pic_email, pic_telepon, created_at, updated_at 
	FROM cabang WHERE id_brand = $1`
	rows, err := r.repo.Db.Query(query, brandID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
		response = append(response, cabang)
	}

	return response, nil
}

func (r cabangPostgresRepository) FindAll() (response []entities.Cabang, err error) {
	query := `SELECT id_cabang, id_brand, nama_cabang, alamat, kota, provinsi, kode_pos, telepon, email, 
		pic_nama, pic_email, pic_telepon, created_at, updated_at 
	FROM cabang`
	rows, err := r.repo.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cabang entities.Cabang
		err = rows.Scan(
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
		response = append(response, cabang)
	}

	return response, nil
}

func (r cabangPostgresRepository) FindAllWithPagination(request entities.DTRequest[entities.Cabang]) (response []entities.Cabang, recordsTotal int, err error) {

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
	countQuery := `SELECT COUNT() FROM cabang`
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
	rows, err := r.repo.Db.Query(baseQuery)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var cabang entities.Cabang
		err = rows.Scan(
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
			return response, 0, err
		}
		response = append(response, cabang)
	}

	// Execute the count query
	err = r.repo.Db.QueryRow(countQuery).Scan(&recordsTotal)
	if err != nil {
		return response, 0, err
	}

	return response, recordsTotal, nil
}

func (r cabangPostgresRepository) Update(cabang entities.Cabang) error {
	query := `UPDATE cabang SET id_brand = $1, nama_cabang = $2, alamat = $3, kota = $4, provinsi = $5, 
		kode_pos = $6, telepon = $7, email = $8, pic_nama = $9, pic_email = $10, pic_telepon = $11, 
		updated_at = NOW() WHERE id_cabang = $12`
	_, err := r.repo.Db.Exec(query, cabang.BrandID, cabang.Name, cabang.Address, cabang.City, cabang.Province,
		cabang.PostalCode, cabang.Phone, cabang.Email, cabang.PICName, cabang.PICEmail, cabang.PICTelepon, cabang.ID)
	return err
}

func (r cabangPostgresRepository) Delete(id int) error {
	query := `DELETE FROM cabang WHERE id_cabang = $1`
	_, err := r.repo.Db.Exec(query, id)
	return err
}
