package entities

import (
	"time"
)

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
	ID         int       `json:"id"`
	BrandID    int       `json:"brand_id"`
	BrandName  int       `json:"nama_brand"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	PostalCode string    `json:"postal_code"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	PICName    string    `json:"pic_name"`
	PICEmail   string    `json:"pic_email"`
	PICTelepon string    `json:"pic_telepon"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Outlet struct {
	ID         int       `json:"id"`
	CabangID   int       `json:"cabang_id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	PostalCode string    `json:"postal_code"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Latitude   *float64  `json:"latitude"`
	Longitude  *float64  `json:"longitude"`
	OpenTime   string    `json:"open_time"`
	CloseTime  string    `json:"close_time"`
	PICName    string    `json:"pic_name"`
	PICEmail   string    `json:"pic_email"`
	PICTelepon string    `json:"pic_telepon"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Transaction struct {
	ID             int        `json:"id"`
	CustomerID     int        `json:"id_pelanggan"`
	OutletID       int        `json:"id_outlet"`
	UserID         *int       `json:"id_access"`
	InvoiceNumber  string     `json:"nomor_invoice"`
	EntryDate      *time.Time `json:"tanggal_masuk"`
	CompletionDate *time.Time `json:"tanggal_selesai"`
	PickupDate     *time.Time `json:"tanggal_diambil"`
	TotalPrice     float64    `json:"total_harga"`
	PaidAmount     float64    `json:"uang_bayar"`
	ChangeAmount   float64    `json:"uang_kembalian"`
	Status         string     `json:"status_transaksi"`
	Note           string     `json:"catatan"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	CreatedBy      *string    `json:"created_by"`
	UpdatedBy      *string    `json:"updated_by"`
}

type TransactionDetail struct {
	ID            int       `json:"id"`
	TransactionID int       `json:"id_transaksi"`
	ServiceID     int       `json:"id_layanan"`
	Quantity      *float64  `json:"kuantitas"`
	Price         *float64  `json:"harga_satuan"`
	Subtotal      *float64  `json:"subtotal"`
	Status        *float64  `json:"status_pengerjaan"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedBy     *string   `json:"created_by"`
	UpdatedBy     *string   `json:"updated_by"`
}

type Employee struct {
	ID        int       `json:"id_pegawai"`
	OutletID  int       `json:"id_outlet"`
	NIK       string    `json:"nik"`
	Name      string    `json:"nama_lengkap"`
	Email     string    `json:"email"`
	Phone     string    `json:"telepon"`
	Address   string    `json:"alamat"`
	BirthDate string    `json:"tanggal_lahir"`
	Gender    string    `json:"jenis_kelamin"`
	Position  string    `json:"posisi"`
	Salary    float64   `json:"gaji"`
	JoinDate  string    `json:"tanggal_masuk"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Customer struct {
	ID        int       `json:"id_pelanggan"`
	OutletID  int       `json:"id_outlet"`
	Name      string    `json:"nama"`
	Email     string    `json:"email"`
	Phone     string    `json:"telepon"`
	Address   string    `json:"alamat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Payment struct {
	ID                   int        `json:"id_pembayaran"`
	TransactionID        int        `json:"id_transaksi"`
	PaymentDate          *time.Time `json:"tanggal_bayar"`
	Amount               float64    `json:"jumlah_bayar"`
	PaymentMethodID      int        `json:"id_metode_pembayaran"`
	Method               string     `json:"metode_bayar"`
	PartnerReferenceNo   string     `json:"nomor_referensi_partner"`
	PartnerStatusCode    string     `json:"status_code_partner"`
	PartnerStatusMessage string     `json:"status_message_partner"`
	Note                 string     `json:"catatan"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}

type HistoryStatusTransaction struct {
	ID            int        `json:"id_history"`
	TransactionID int        `json:"id_transaksi"`
	OldStatus     string     `json:"status_lama"`
	NewStatus     string     `json:"status_baru"`
	ChangeTime    *time.Time `json:"waktu_perubahan"`
	Description   string     `json:"keterangan"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
