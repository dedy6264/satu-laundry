package cabangrepo

import (
	"laundry-backend/entities"
)

type CabangRepo interface {
	Create(cabang entities.Cabang) error
	FindByID(id int) (response entities.Cabang, err error)
	FindByBrandID(brandID int) (response []entities.Cabang, err error)
	FindAll() (response []entities.Cabang, err error)
	FindAllWithPagination(request entities.DTRequest[entities.Cabang]) (response []entities.Cabang, recordsTotal int, err error)
	Update(cabang entities.Cabang) error
	Delete(id int) error
}
