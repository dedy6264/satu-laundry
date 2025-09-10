package repositories

import (
	"database/sql"
	"laundry-backend/internal/entities"
	"fmt"
	"strings"
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

func (r *brandPostgresRepository) FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Brand, int, int, error) {
	// Validate orderBy field to prevent SQL injection
	validFields := map[string]bool{
		"id":          true,
		"name":        true,
		"description": true,
		"pic_name":    true,
		"pic_email":   true,
		"created_at":  true,
		"updated_at":  true,
	}
	
	// Map field names to database column names
	fieldMap := map[string]string{
		"id":          "id_brand",
		"name":        "nama_brand",
		"description": "deskripsi",
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
	baseQuery := `SELECT id_brand, nama_brand, deskripsi, pic_nama, pic_email, pic_telepon, logo_url, created_at, updated_at FROM brand`
	countQuery := `SELECT COUNT(*) FROM brand`
	
	var args []interface{}
	argIndex := 1
	
	// Add search condition if provided
	if search != "" {
		search = strings.ToLower(search)
		baseQuery += fmt.Sprintf(` WHERE LOWER(nama_brand) LIKE $%d OR LOWER(pic_nama) LIKE $%d OR LOWER(pic_email) LIKE $%d`, argIndex, argIndex+1, argIndex+2)
		countQuery += fmt.Sprintf(` WHERE LOWER(nama_brand) LIKE $%d OR LOWER(pic_nama) LIKE $%d OR LOWER(pic_email) LIKE $%d`, argIndex, argIndex+1, argIndex+2)
		args = append(args, "%"+search+"%", "%"+search+"%", "%"+search+"%")
		argIndex += 3
	}
	
	// Add ordering
	dbOrderBy := fieldMap[orderBy]
	if dbOrderBy == "" {
		dbOrderBy = "id_brand"
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
			return nil, 0, 0, err
		}
		brands = append(brands, brand)
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
	
	return brands, recordsTotal, recordsFiltered, nil
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