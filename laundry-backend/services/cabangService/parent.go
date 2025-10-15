package administrationservice

import (
	"laundry-backend/services"
)

type CabangService struct {
	service services.UsecaseService
}

func ApiCabangService(service services.UsecaseService) CabangService {
	return CabangService{
		service: service,
	}
}