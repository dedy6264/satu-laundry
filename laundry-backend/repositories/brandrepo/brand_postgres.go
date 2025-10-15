package brandrepo

import (
	"database/sql"
	"laundry-backend/entities"
	"laundry-backend/repositories"
	"strconv"
)

type brandPostgresRepository struct {
	repo repositories.Repositories
}

func NewBrandRepo(repo repositories.Repositories) brandPostgresRepository {
	return brandPostgresRepository{repo: repo}
}

func (r brandPostgresRepository) Create(brand entities.Brand) error {
	query := `INSERT INTO brand (nama_brand, deskripsi, pic_nama, pic_email, pic_telepon, logo_url, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW()) RETURNING id_brand`
	return r.repo.Db.QueryRow(query, brand.Name, brand.Description, brand.PICName, brand.PICEmail, brand.PICTelepon, brand.LogoURL).Scan(&brand.ID)
}

func (r brandPostgresRepository) FindByID(id int) (response entities.Brand, err error) {
	query := `SELECT id_brand, nama_brand, deskripsi, pic_nama, pic_email, pic_telepon, logo_url, created_at, updated_at 
	FROM brand WHERE id_brand = $1`
	row := r.repo.Db.QueryRow(query, id)

	err = row.Scan(
		&response.ID,
		&response.Name,
		&response.Description,
		&response.PICName,
		&response.PICEmail,
		&response.PICTelepon,
		&response.LogoURL,
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

func (r brandPostgresRepository) FindAll() (response []entities.Brand, err error) {
	query := `SELECT id_brand, nama_brand, deskripsi, pic_nama, pic_email, pic_telepon, logo_url, created_at, updated_at FROM brand`
	rows, err := r.repo.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var brand entities.Brand
		err := rows.Scan(
			&brand.ID,
			&brand.Name,
			&brand.Description,
			&brand.PICName,
			&brand.PICEmail,
			&brand.PICTelepon,
			&brand.LogoURL,
			&brand.CreatedAt,
			&brand.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		response = append(response, brand)
	}

	return response, nil
}

func (r brandPostgresRepository) FindAllWithPagination(request entities.DTRequest[entities.Brand]) (response []entities.Brand, count int, err error) {

	// Build the query
	getQuery := `SELECT id_brand, nama_brand, deskripsi, pic_nama, pic_email, pic_telepon, logo_url, created_at, updated_at FROM brand where true `
	countQuery := `SELECT COUNT() FROM brand`
	if request.Data.ID != 0 {
		getQuery += ` and.id_brand = ` + strconv.Itoa(request.Data.ID)
	}
	if request.OrderBy != "" {
		getQuery += ` ORDER BY ` + request.OrderBy + ` ` + request.SortBy
	} else {
		getQuery += ` ORDER BY id_transaksi ASC`
	}
	if request.Length != 0 {
		getQuery += ` LIMIT ` + strconv.Itoa(request.Length) + ` OFFSET ` + strconv.Itoa(request.Start)
	}
	// Execute the data query
	rows, err := r.repo.Db.Query(getQuery)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var brand entities.Brand
		err := rows.Scan(
			&brand.ID,
			&brand.Name,
			&brand.Description,
			&brand.PICName,
			&brand.PICEmail,
			&brand.PICTelepon,
			&brand.LogoURL,
			&brand.CreatedAt,
			&brand.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		response = append(response, brand)
	}

	// Get total count

	err = r.repo.Db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return response, count, nil
}

func (r brandPostgresRepository) Update(brand entities.Brand) error {
	query := `UPDATE brand SET nama_brand = $1, deskripsi = $2, pic_nama = $3, pic_email = $4, pic_telepon = $5, 
	logo_url = $6, updated_at = NOW() WHERE id_brand = $7`
	_, err := r.repo.Db.Exec(query, brand.Name, brand.Description, brand.PICName, brand.PICEmail, brand.PICTelepon, brand.LogoURL, brand.ID)
	return err
}

func (r brandPostgresRepository) Delete(id int) error {
	query := `DELETE FROM brand WHERE id_brand = $1`
	_, err := r.repo.Db.Exec(query, id)
	return err
}
