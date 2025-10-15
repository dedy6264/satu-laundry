package repositories

import (
	"database/sql"
	"laundry-backend/internal/entities"
	"strconv"
)

type brandPostgresRepository struct {
	db *sql.DB
}

func NewBrandRepository(db *sql.DB) BrandRepository {
	return &brandPostgresRepository{db: db}
}

func (r *brandPostgresRepository) Create(brand *entities.Brand) error {
	query := `INSERT INTO brand (nama_brand, deskripsi, pic_nama, pic_email, pic_telepon, logo_url, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW()) RETURNING id_brand`
	return r.db.QueryRow(query, brand.Name, brand.Description, brand.PICName, brand.PICEmail, brand.PICTelepon, brand.LogoURL).Scan(&brand.ID)
}

func (r *brandPostgresRepository) FindByID(id int) (*entities.Brand, error) {
	query := `SELECT id_brand, nama_brand, deskripsi, pic_nama, pic_email, pic_telepon, logo_url, created_at, updated_at 
	FROM brand WHERE id_brand = $1`
	row := r.db.QueryRow(query, id)

	var brand entities.Brand
	err := row.Scan(
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
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &brand, nil
}

func (r *brandPostgresRepository) FindAll() ([]entities.Brand, error) {
	query := `SELECT id_brand, nama_brand, deskripsi, pic_nama, pic_email, pic_telepon, logo_url, created_at, updated_at FROM brand`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brands []entities.Brand
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
		brands = append(brands, brand)
	}

	return brands, nil
}

func (r *brandPostgresRepository) FindAllWithPagination(request entities.DTRequest[entities.Brand]) ([]entities.Brand, int, error) {

	// Build the query
	getQuery := `SELECT id_brand, nama_brand, deskripsi, pic_nama, pic_email, pic_telepon, logo_url, created_at, updated_at FROM brand where true `
	countQuery := `SELECT COUNT(*) FROM brand`
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
	rows, err := r.db.Query(getQuery)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var brands []entities.Brand
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
		brands = append(brands, brand)
	}

	// Get total count
	var totalCount int
	err = r.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	return brands, totalCount, nil
}

func (r *brandPostgresRepository) Update(brand *entities.Brand) error {
	query := `UPDATE brand SET nama_brand = $1, deskripsi = $2, pic_nama = $3, pic_email = $4, pic_telepon = $5, 
	logo_url = $6, updated_at = NOW() WHERE id_brand = $7`
	_, err := r.db.Exec(query, brand.Name, brand.Description, brand.PICName, brand.PICEmail, brand.PICTelepon, brand.LogoURL, brand.ID)
	return err
}

func (r *brandPostgresRepository) Delete(id int) error {
	query := `DELETE FROM brand WHERE id_brand = $1`
	_, err := r.db.Exec(query, id)
	return err
}
