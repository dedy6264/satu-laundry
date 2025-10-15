package servicecategoryrepo

import (
	"laundry-backend/entities"
)

type ServiceCategoryRepo interface {
	Create(category entities.ServiceCategory) error
	FindByID(id int) (entities.ServiceCategory, error)
	FindAll() ([]entities.ServiceCategory, error)
	FindAllWithPagination(request entities.DTRequest[entities.ServiceCategory]) ([]entities.ServiceCategory, int, error)
	Update(category entities.ServiceCategory) error
	Delete(id int) error
}
