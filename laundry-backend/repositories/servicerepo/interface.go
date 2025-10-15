package servicerepo

import (
	"laundry-backend/entities"
)

type ServiceRepo interface {
	Create(service entities.Service) error
	FindByID(id int) (entities.Service, error)
	FindAll(request entities.Service) ([]entities.Service, error)
	FindAllWithPagination(request entities.DTRequest[entities.Service]) ([]entities.Service, int, error)
	Update(service entities.Service) error
	Delete(id int) error
	FindByCategoryID(categoryID int) ([]entities.Service, error)
}
