package customerservice

import (
	"laundry-backend/services"
)

type CustomerService struct {
	service services.UsecaseService
}

func ApiCustomerService(service services.UsecaseService) CustomerService {
	return CustomerService{
		service: service,
	}
}
