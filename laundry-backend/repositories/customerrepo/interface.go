package customerrepo

import (
	"laundry-backend/entities"
)

type CustomerRepo interface {
	Create(customer entities.Customer) error
	FindByID(id int) (entities.Customer, error)
	FindByOutletID(outletID int) ([]entities.Customer, error)
	FindAll() ([]entities.Customer, error)
	FindAllWithPagination(request entities.DTRequest[entities.Customer]) ([]entities.Customer, int, error)
	Update(customer entities.Customer) error
	Delete(id int) error
}
