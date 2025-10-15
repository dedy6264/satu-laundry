package administrationservice

import (
	"laundry-backend/services"
)

type BrandService struct {
	service services.UsecaseService
}

func ApiBrandService(service services.UsecaseService) BrandService {
	return BrandService{
		service: service,
	}
}