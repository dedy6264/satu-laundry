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

type InquiryRequest struct {
	ServicePackageID int     `json:"id_layanan" validare:"required"`
	CustomerID       int     `json:"id_pelanggan" validare:"required"`
	OutletID         int     `json:"id_outlet"`
	EmployeeID       int     `json:"id_pegawai"`
	Quantity         float64 `json:"jumlah" validare:"required"`
	Note             string  `json:"catatan"`
}

type RegisterEmployeeRequest struct {
	OutletID  int     `json:"id_outlet"`
	NIK       string  `json:"nik"`
	Name      string  `json:"nama_lengkap"`
	Email     string  `json:"email"`
	Phone     string  `json:"telepon"`
	Address   string  `json:"alamat"`
	BirthDate string  `json:"tanggal_lahir"`
	Gender    string  `json:"jenis_kelamin"`
	Position  string  `json:"posisi"`
	Salary    float64 `json:"gaji"`
	JoinDate  string  `json:"tanggal_masuk"`
	Status    string  `json:"status"`
	Password  string  `json:"password"`
}

type EmployeeLoginRequest struct {
	Email    string `json:"email"` // Can be email, NIK, or phone
	Password string `json:"password"`
}

type RegisterCustomerRequest struct {
	OutletID int    `json:"id_outlet"`
	Name     string `json:"nama"`
	Email    string `json:"email"`
	Phone    string `json:"telepon"`
	Address  string `json:"alamat"`
}
