package repositories

import (
	"laundry-backend/internal/entities"
)

type PaymentMethodRepository interface {
	Create(paymentMethod *entities.PaymentMethod) error
	FindByID(id int) (*entities.PaymentMethod, error)
	FindAll() ([]entities.PaymentMethod, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.PaymentMethod, int, int, error)
	Update(paymentMethod *entities.PaymentMethod) error
	Delete(id int) error
}