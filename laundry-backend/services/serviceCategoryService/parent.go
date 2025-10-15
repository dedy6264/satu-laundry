package administrationservice

import (
	"laundry-backend/services"
)

type ServiceCategoryService struct {
	service services.UsecaseService
}

func ApiServiceCategoryService(service services.UsecaseService) ServiceCategoryService {
	return ServiceCategoryService{
		service: service,
	}
}