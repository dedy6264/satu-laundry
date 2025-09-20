package entities

import (
	"time"
)

type PaymentMethod struct {
	ID          int       `json:"id" gorm:"primaryKey;column:id"`
	NamaMetode  string    `json:"nama_metode" gorm:"column:nama_metode"`
	URL         string    `json:"url" gorm:"column:url"`
	SKey        string    `json:"s_key" gorm:"column:s_key"`
	MKey        string    `json:"m_key" gorm:"column:m_key"`
	MerchantFee float64   `json:"merchant_fee" gorm:"column:merchant_fee"`
	AdminFee    float64   `json:"admin_fee" gorm:"column:admin_fee"`
	Status      string    `json:"status" gorm:"column:status"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy   string    `json:"created_by" gorm:"column:created_by"`
	UpdatedBy   string    `json:"updated_by" gorm:"column:updated_by"`
}

func (PaymentMethod) TableName() string {
	return "metode_pembayaran"
}
