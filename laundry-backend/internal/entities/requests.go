package entities

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type RegisterBrandRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	PICName     string `json:"pic_name"`
	PICEmail    string `json:"pic_email"`
	PICTelepon  string `json:"pic_telepon"`
	LogoURL     string `json:"logo_url"`
}

type RegisterCabangRequest struct {
	BrandID    int    `json:"brand_id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postal_code"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	PICName    string `json:"pic_name"`
	PICEmail   string `json:"pic_email"`
	PICTelepon string `json:"pic_telepon"`
}

type RegisterOutletRequest struct {
	CabangID   int     `json:"cabang_id"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	City       string  `json:"city"`
	Province   string  `json:"province"`
	PostalCode string  `json:"postal_code"`
	Phone      string  `json:"phone"`
	Email      string  `json:"email"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	OpenTime   string  `json:"open_time"`
	CloseTime  string  `json:"close_time"`
	PICName    string  `json:"pic_name"`
	PICEmail   string  `json:"pic_email"`
	PICTelepon string  `json:"pic_telepon"`
}