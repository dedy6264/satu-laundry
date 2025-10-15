package employeerepo

import (
	"laundry-backend/entities"
)

type EmployeeRepo interface {
	Create(employee entities.Employee) error
	FindByID(id int) (entities.Employee, error)
	FindAll() ([]entities.Employee, error)
	FindAllWithPagination(request entities.DTRequest[entities.Employee]) ([]entities.Employee, int, error)
	Update(employee entities.Employee) error
	Delete(id int) error
}
