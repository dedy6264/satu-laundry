package administrationservice

import (
	"laundry-backend/services"
)

type EmployeeService struct {
	service services.UsecaseService
}

func ApiEmployeeService(service services.UsecaseService) EmployeeService {
	return EmployeeService{
		service: service,
	}
}