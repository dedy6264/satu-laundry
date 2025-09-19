package entities

import (
	"time"
)

type Service struct {
	BrandID     int       `json:"brand_id" `
	ID          int       `json:"id"`


	CategoryID  int       `json:"kategori_id"`
	Name        string    `json:"nama_layanan"`
	Description string    `json:"deskripsi"`
	Price       float64   `json:"harga_satuan"`
	Unit        string    `json:"satuan_durasi"`
	Estimation  int       `json:"durasi_pengerjaan"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ServiceCategory struct {
	ID          int       `json:"id"`
	Name        string    `json:"nama_kategori"`
	Description string    `json:"deskripsi"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateServiceRequest struct {
	BrandID     int     `json:"brand_id" validare:"required"`
	CategoryID  int     `json:"kategori_id" validare:"required"`
	Name        string  `json:"nama_layanan" validare:"required"`
	Description string  `json:"deskripsi"`
	Price       float64 `json:"harga_satuan" validare:"required"`
	Unit        string  `json:"satuan_durasi" validare:"required"`
	Estimation  int     `json:"durasi_pengerjaan"`
}

type UpdateServiceRequest struct {
	CategoryID  int     `json:"kategori_id" validare:"required"`
	Name        string  `json:"nama_layanan" validare:"required"`
	Description string  `json:"deskripsi"`
	Price       float64 `json:"harga_satuan" validare:"required"`
	Unit        string  `json:"satuan_durasi" validare:"required"`
	Estimation  int     `json:"durasi_pengerjaan"`
}

type CreateServiceCategoryRequest struct {
	Name        string `json:"nama_kategori" validare:"required"`
	Description string `json:"deskripsi"`
}

type UpdateServiceCategoryRequest struct {
	Name        string `json:"nama_kategori" validare:"required"`
	Description string `json:"deskripsi"`
}
