package outletrepo

import (
	"laundry-backend/entities"
)

type OutletRepo interface {
	Create(outlet entities.Outlet) error
	FindByID(id int) (entities.Outlet, error)
	FindByCabangID(cabangID int) ([]entities.Outlet, error)
	FindAll(request entities.Outlet) ([]entities.Outlet, error)
	FindAllWithPagination(request entities.DTRequest[entities.Outlet]) ([]entities.Outlet, int, error)
	Update(outlet entities.Outlet) error
	Delete(id int) error
}
