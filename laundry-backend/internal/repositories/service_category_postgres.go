package repositories

import (
	"database/sql"
	"laundry-backend/internal/entities"
	"strconv"
)

type serviceCategoryPostgresRepository struct {
	db *sql.DB
}

func NewServiceCategoryRepository(db *sql.DB) ServiceCategoryRepository {
	return &serviceCategoryPostgresRepository{
		db: db,
	}
}

func (r *serviceCategoryPostgresRepository) Create(category *entities.ServiceCategory) error {
	query := `
		INSERT INTO kategori_layanan (nama_kategori, deskripsi, created_at, updated_at)
		VALUES ($1, $2, NOW(), NOW())
		RETURNING id_kategori`

	err := r.db.QueryRow(query, category.Name, category.Description).
		Scan(&category.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *serviceCategoryPostgresRepository) FindByID(id int) (*entities.ServiceCategory, error) {
	query := `
		SELECT id_kategori, nama_kategori, deskripsi, created_at, updated_at
		FROM kategori_layanan
		WHERE id_kategori = $1`

	var category entities.ServiceCategory
	err := r.db.QueryRow(query, id).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &category, nil
}

func (r *serviceCategoryPostgresRepository) FindAll() ([]entities.ServiceCategory, error) {
	query := `
		SELECT id_kategori, nama_kategori, deskripsi, created_at, updated_at
		FROM kategori_layanan
		ORDER BY id_kategori`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entities.ServiceCategory
	for rows.Next() {
		var category entities.ServiceCategory
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *serviceCategoryPostgresRepository) FindAllWithPagination(request entities.DTRequest[entities.ServiceCategory]) ([]entities.ServiceCategory, int, error) {
	// Base query
	baseQuery := `
		FROM kategori_layanan where true `

	// Count query
	countQuery := "SELECT COUNT(*) " + baseQuery

	// Data query
	dataQuery := `
		SELECT id_kategori, nama_kategori, deskripsi, created_at, updated_at
		` + baseQuery
	if request.Data.ID != 0 {
		baseQuery += ` and id_kategori = ` + strconv.Itoa(request.Data.ID)
	}
	if request.OrderBy != "" {
		baseQuery += ` ORDER BY ` + request.OrderBy + ` ` + request.SortBy
	} else {
		baseQuery += ` ORDER BY id_kategori ASC`
	}
	if request.Length != 0 {
		baseQuery += ` LIMIT ` + strconv.Itoa(request.Length) + ` OFFSET ` + strconv.Itoa(request.Start)
	}
	// Get total count
	var totalCount int
	err := r.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	// Execute data query
	rows, err := r.db.Query(dataQuery)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var categories []entities.ServiceCategory
	for rows.Next() {
		var category entities.ServiceCategory
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		categories = append(categories, category)
	}

	// Return categories, total count, filtered count, and error
	return categories, totalCount, nil
}

func (r *serviceCategoryPostgresRepository) Update(category *entities.ServiceCategory) error {
	query := `
		UPDATE kategori_layanan
		SET nama_kategori = $1, deskripsi = $2, updated_at = NOW()
		WHERE id_kategori = $3`

	_, err := r.db.Exec(query, category.Name, category.Description, category.ID)
	return err
}

func (r *serviceCategoryPostgresRepository) Delete(id int) error {
	query := `DELETE FROM kategori_layanan WHERE id_kategori = $1`
	_, err := r.db.Exec(query, id)
	return err
}
