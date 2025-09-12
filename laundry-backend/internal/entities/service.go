package entities

import (
	"time"
)

type Service struct {
	ID          int       `json:"id"`
	CategoryID  int       `json:"id_kategori"`
	Name        string    `json:"nama_layanan"`
	Description string    `json:"deskripsi"`
	Price       float64   `json:"harga"`
	Unit        string    `json:"satuan"`
	Estimation  string    `json:"estimasi_waktu"`
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
	CategoryID  int     `json:"id_kategori" validare:"required"`
	Name        string  `json:"nama_layanan" validare:"required"`
	Description string  `json:"deskripsi"`
	Price       float64 `json:"harga" validare:"required"`
	Unit        string  `json:"satuan" validare:"required"`
	Estimation  string  `json:"estimasi_waktu"`
}

type UpdateServiceRequest struct {
	CategoryID  int     `json:"id_kategori" validare:"required"`
	Name        string  `json:"nama_layanan" validare:"required"`
	Description string  `json:"deskripsi"`
	Price       float64 `json:"harga" validare:"required"`
	Unit        string  `json:"satuan" validare:"required"`
	Estimation  string  `json:"estimasi_waktu"`
}

type CreateServiceCategoryRequest struct {
	Name        string `json:"nama_kategori" validare:"required"`
	Description string `json:"deskripsi"`
}

type UpdateServiceCategoryRequest struct {
	Name        string `json:"nama_kategori" validare:"required"`
	Description string `json:"deskripsi"`
}