package serviceservice

import (
	"laundry-backend/services"
)

type ServiceService struct {
	service services.UsecaseService
}

func ApiServiceService(service services.UsecaseService) ServiceService {
	return ServiceService{
		service: service,
	}
}
