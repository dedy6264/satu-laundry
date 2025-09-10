package entities

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Brand struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PICName     string    `json:"pic_name"`
	PICEmail    string    `json:"pic_email"`
	PICTelepon  string    `json:"pic_telepon"`
	LogoURL     string    `json:"logo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Cabang struct {
	ID          int       `json:"id"`
	BrandID     int       `json:"brand_id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	PostalCode  string    `json:"postal_code"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	PICName     string    `json:"pic_name"`
	PICEmail    string    `json:"pic_email"`
	PICTelepon  string    `json:"pic_telepon"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Outlet struct {
	ID          int       `json:"id"`
	CabangID    int       `json:"cabang_id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	PostalCode  string    `json:"postal_code"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	OpenTime    string    `json:"open_time"`
	CloseTime   string    `json:"close_time"`
	PICName     string    `json:"pic_name"`
	PICEmail    string    `json:"pic_email"`
	PICTelepon  string    `json:"pic_telepon"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}