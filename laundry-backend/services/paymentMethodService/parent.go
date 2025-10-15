package paymentmethodservice

import (
	"laundry-backend/services"
)

type PaymentMethodService struct {
	service services.UsecaseService
}

func ApiPaymentMethodService(service services.UsecaseService) PaymentMethodService {
	return PaymentMethodService{
		service: service,
	}
}
