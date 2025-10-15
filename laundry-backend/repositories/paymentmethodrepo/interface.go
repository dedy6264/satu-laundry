package paymentmethodrepo

import (
	"laundry-backend/entities"
)

type PaymentMethodRepo interface {
	Create(paymentMethod entities.PaymentMethod) error
	FindByID(id int) (entities.PaymentMethod, error)
	FindAll() ([]entities.PaymentMethod, error)
	FindAllWithPagination(request entities.DTRequest[entities.PaymentMethod]) ([]entities.PaymentMethod, int, error)
	Update(paymentMethod entities.PaymentMethod) error
	Delete(id int) error
}
