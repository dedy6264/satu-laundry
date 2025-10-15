package administrationservice

import (
	"laundry-backend/services"
)

type OutletService struct {
	service services.UsecaseService
}

func ApiOutletService(service services.UsecaseService) OutletService {
	return OutletService{
		service: service,
	}
}