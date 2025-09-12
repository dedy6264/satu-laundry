package repositories

import (
	"database/sql"
	"fmt"
	"laundry-backend/internal/entities"
)

type servicePostgresRepository struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) ServiceRepository {
	return &servicePostgresRepository{
		db: db,
	}
}

func (r *servicePostgresRepository) Create(service *entities.Service) error {
	query := `
		INSERT INTO paket_layanan (id_brand, id_kategori, nama_layanan, deskripsi, harga_per_kg, satuan_durasi, durasi_pengerjaan, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
		RETURNING id_layanan`

	err := r.db.QueryRow(query, service.BrandID, service.CategoryID, service.Name, service.Description, service.Price, service.Unit, service.Estimation).
		Scan(&service.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *servicePostgresRepository) FindByID(id int) (*entities.Service, error) {
	query := `
		SELECT l.id_layanan, l.id_brand, l.id_kategori, l.nama_layanan, l.deskripsi, l.harga_per_kg, l.satuan_durasi, l.durasi_pengerjaan, l.created_at, l.updated_at
		FROM paket_layanan l
		WHERE l.id_layanan = $1`

	var service entities.Service
	err := r.db.QueryRow(query, id).Scan(
		&service.ID,
		&service.BrandID,
		&service.CategoryID,
		&service.Name,
		&service.Description,
		&service.Price,
		&service.Unit,
		&service.Estimation,
		&service.CreatedAt,
		&service.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &service, nil
}

func (r *servicePostgresRepository) FindAll() ([]entities.Service, error) {
	query := `
		SELECT l.id_layanan, l.id_brand, l.id_kategori, l.nama_layanan, l.deskripsi, l.harga_per_kg, l.satuan_durasi, l.durasi_pengerjaan, l.created_at, l.updated_at
		FROM paket_layanan l
		ORDER BY l.id_layanan`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []entities.Service
	for rows.Next() {
		var service entities.Service
		err := rows.Scan(
			&service.ID,
			&service.BrandID,
			&service.CategoryID,
			&service.Name,
			&service.Description,
			&service.Price,
			&service.Unit,
			&service.Estimation,
			&service.CreatedAt,
			&service.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	return services, nil
}

func (r *servicePostgresRepository) FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Service, int, error) {
	// Base query
	baseQuery := `
		FROM paket_layanan l`

	// Count query
	countQuery := "SELECT COUNT(*) " + baseQuery

	// Data query
	dataQuery := `
		SELECT l.id_layanan, l.id_brand, l.id_kategori, l.nama_layanan, l.deskripsi, l.harga_per_kg, l.satuan_durasi, l.durasi_pengerjaan, l.created_at, l.updated_at
		` + baseQuery

	// Search condition
	var args []interface{}
	if search != "" {
		countQuery += " WHERE l.nama_layanan ILIKE $1"
		dataQuery += " WHERE l.nama_layanan ILIKE $1"
		args = append(args, "%"+search+"%")
	}

	// Get total count
	var totalCount int
	countArgs := make([]interface{}, len(args))
	copy(countArgs, args)
	err := r.db.QueryRow(countQuery, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	// Add ordering
	if orderBy == "" {
		orderBy = "l.id_layanan"
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
		return nil, 0, err
	}
	defer rows.Close()

	var services []entities.Service
	for rows.Next() {
		var service entities.Service
		err := rows.Scan(
			&service.ID,
			&service.BrandID,
			&service.CategoryID,
			&service.Name,
			&service.Description,
			&service.Price,
			&service.Unit,
			&service.Estimation,
			&service.CreatedAt,
			&service.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		services = append(services, service)
	}

	return services, totalCount, nil
}

func (r *servicePostgresRepository) Update(service *entities.Service) error {
	query := `
		UPDATE paket_layanan
		SET id_brand = $1, id_kategori = $2, nama_layanan = $3, deskripsi = $4, harga_per_kg = $5, satuan_durasi = $6, durasi_pengerjaan = $7, updated_at = NOW()
		WHERE id_layanan = $8`

	_, err := r.db.Exec(query, service.BrandID, service.CategoryID, service.Name, service.Description, service.Price, service.Unit, service.Estimation, service.ID)
	return err
}

func (r *servicePostgresRepository) Delete(id int) error {
	query := `DELETE FROM paket_layanan WHERE id_layanan = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *servicePostgresRepository) FindByCategoryID(categoryID int) ([]entities.Service, error) {
	query := `
		SELECT l.id_layanan, l.id_brand, l.id_kategori, l.nama_layanan, l.deskripsi, l.harga_per_kg, l.satuan_durasi, l.durasi_pengerjaan, l.created_at, l.updated_at
		FROM paket_layanan l
		WHERE l.id_kategori = $1
		ORDER BY l.id_layanan`

	rows, err := r.db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []entities.Service
	for rows.Next() {
		var service entities.Service
		err := rows.Scan(
			&service.ID,
			&service.BrandID,
			&service.CategoryID,
			&service.Name,
			&service.Description,
			&service.Price,
			&service.Unit,
			&service.Estimation,
			&service.CreatedAt,
			&service.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	return services, nil
}
