package repositories

import (
	"database/sql"
	"laundry-backend/internal/entities"
	"strconv"
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
		INSERT INTO paket_layanan (id_brand, id_kategori, nama_layanan, deskripsi,satuan, harga_satuan, satuan_durasi, durasi_pengerjaan, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7,$8, NOW(), NOW())
		RETURNING id_layanan`

	err := r.db.QueryRow(query, service.BrandID, service.CategoryID, service.Name, service.Description, service.Satuan, service.Price, service.Unit, service.Estimation).
		Scan(&service.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *servicePostgresRepository) FindByID(id int) (*entities.Service, error) {
	query := `
		SELECT l.id_layanan, l.id_brand, l.id_kategori, l.nama_layanan, l.deskripsi,l.satuan, l.harga_satuan, l.satuan_durasi, l.durasi_pengerjaan, l.created_at, l.updated_at
		FROM paket_layanan l
		WHERE l.id_layanan = $1`

	var service entities.Service
	err := r.db.QueryRow(query, id).Scan(
		&service.ID,
		&service.BrandID,
		&service.CategoryID,
		&service.Name,
		&service.Description,
		&service.Satuan,
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

func (r *servicePostgresRepository) FindAll(request entities.Service) ([]entities.Service, error) {
	query := `
		SELECT id_layanan, id_brand, id_kategori, nama_layanan, deskripsi,satuan, harga_satuan, satuan_durasi, durasi_pengerjaan, created_at, updated_at
		FROM paket_layanan l where true 
		`
	if request.ID != 0 {
		query += ` and id_layanan = ` + strconv.Itoa(request.ID)
	}
	if request.BrandID != 0 {
		query += ` and id_brand = ` + strconv.Itoa(request.BrandID)
	}
	if request.CategoryID != 0 {
		query += ` and id_kategori = ` + strconv.Itoa(request.CategoryID)
	}
	query += ` ORDER BY id_layanan`
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
			&service.Satuan,
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
	if len(services) == 0 {

		return nil, sql.ErrNoRows
	}
	return services, nil
}

func (r *servicePostgresRepository) FindAllWithPagination(request entities.DTRequest[entities.Service]) ([]entities.Service, int, error) {
	// Base query
	baseQuery := ` FROM paket_layanan l where true `
	getQuery := ` SELECT l.id_layanan, l.id_brand, l.id_kategori, l.nama_layanan, l.deskripsi,l.satuan, l.harga_satuan, l.satuan_durasi, l.durasi_pengerjaan, l.created_at, l.updated_at  ` + baseQuery
	countQuery := `SELECT COUNT(*) ` + baseQuery
	if request.Data.BrandID != 0 {
		getQuery += ` and id_brand = ` + strconv.Itoa(request.Data.BrandID)
	}
	if request.Data.CategoryID != 0 {
		getQuery += ` and id_kategori = ` + strconv.Itoa(request.Data.CategoryID)
	}
	if request.OrderBy != "" {
		getQuery += ` ORDER BY ` + request.OrderBy + ` ` + request.SortBy
	} else {
		getQuery += ` ORDER BY l.id_layanan ASC`
	}
	if request.Length != 0 {
		getQuery += ` LIMIT ` + strconv.Itoa(request.Length) + ` OFFSET ` + strconv.Itoa(request.Start)
	}

	// Get total count
	var totalCount int
	err := r.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	// Execute data query
	rows, err := r.db.Query(getQuery)
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
			&service.Satuan,
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
		SET id_brand = $1, id_kategori = $2, nama_layanan = $3, deskripsi = $4, harga_satuan = $5, satuan_durasi = $6, durasi_pengerjaan = $7, updated_at = NOW()
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
		SELECT l.id_layanan, l.id_brand, l.id_kategori, l.nama_layanan, l.deskripsi,l.satuan, l.harga_satuan, l.satuan_durasi, l.durasi_pengerjaan, l.created_at, l.updated_at
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
			&service.Satuan,
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
