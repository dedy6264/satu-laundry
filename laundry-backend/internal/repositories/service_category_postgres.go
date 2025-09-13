package repositories

import (
	"database/sql"
	"fmt"
	"laundry-backend/internal/entities"
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

func (r *serviceCategoryPostgresRepository) FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.ServiceCategory, int, int, error) {
	// Base query
	baseQuery := `
		FROM kategori_layanan k`

	// Count query
	countQuery := "SELECT COUNT(*) " + baseQuery

	// Data query
	dataQuery := `
		SELECT k.id_kategori, k.nama_kategori, k.deskripsi, k.created_at, k.updated_at
		` + baseQuery

	// Search condition
	var args []interface{}
	if search != "" {
		countQuery += " WHERE k.nama_kategori ILIKE $1"
		dataQuery += " WHERE k.nama_kategori ILIKE $1"
		args = append(args, "%"+search+"%")
	}

	// Get total count
	var totalCount int
	countArgs := make([]interface{}, len(args))
	copy(countArgs, args)
	err := r.db.QueryRow(countQuery, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, 0, 0, err
	}

	// Add ordering
	if orderBy == "" {
		orderBy = "k.id_kategori"
	}
	if orderDir == "" {
		orderDir = "ASC"
	}
	dataQuery += fmt.Sprintf(" ORDER BY %s %s LIMIT $%d OFFSET $%d", orderBy, orderDir, len(args)+1, len(args)+2)

	// Add limit and offset
	args = append(args, limit, offset)

	// Execute data query
	rows, err := r.db.Query(dataQuery, args...)
	if err != nil {
		return nil, 0, 0, err
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
			return nil, 0, 0, err
		}
		categories = append(categories, category)
	}

	// Return categories, total count, filtered count, and error
	return categories, totalCount, totalCount, nil
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
