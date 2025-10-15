package brandrepo

import (
	"laundry-backend/entities"
)

type BrandRepo interface {
	FindByID(id int) (response entities.Brand, err error)
	FindAll() (response []entities.Brand, err error)
	FindAllWithPagination(request entities.DTRequest[entities.Brand]) (response []entities.Brand, count int, err error)
	Update(brand entities.Brand) error
	Delete(id int) error
}
