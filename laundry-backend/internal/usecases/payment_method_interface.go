package usecases

import (
	"laundry-backend/internal/entities"
)

type PaymentMethodUsecase interface {
	CreatePaymentMethod(request entities.CreatePaymentMethodRequest) error
	GetPaymentMethodByID(id int) (*entities.PaymentMethod, error)
	GetAllPaymentMethods() ([]entities.PaymentMethod, error)
	GetAllPaymentMethodsDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error)
	UpdatePaymentMethod(id int, request entities.UpdatePaymentMethodRequest) error
	DeletePaymentMethod(id int) error
}