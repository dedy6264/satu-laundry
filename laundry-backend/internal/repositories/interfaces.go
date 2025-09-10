package repositories

import (
	"laundry-backend/internal/entities"
)

type UserRepository interface {
	FindByEmail(email string) (*entities.User, error)
	Create(user *entities.User) error
}

type BrandRepository interface {
	Create(brand *entities.Brand) error
	FindByID(id int) (*entities.Brand, error)
	FindAll() ([]entities.Brand, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Brand, int, int, error)
	Update(brand *entities.Brand) error
	Delete(id int) error
}

type CabangRepository interface {
	Create(cabang *entities.Cabang) error
	FindByID(id int) (*entities.Cabang, error)
	FindByBrandID(brandID int) ([]entities.Cabang, error)
	FindAll() ([]entities.Cabang, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Cabang, int, int, error)
	Update(cabang *entities.Cabang) error
	Delete(id int) error
}

type OutletRepository interface {
	Create(outlet *entities.Outlet) error
	FindByID(id int) (*entities.Outlet, error)
	FindByCabangID(cabangID int) ([]entities.Outlet, error)
	FindAll() ([]entities.Outlet, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Outlet, int, int, error)
	Update(outlet *entities.Outlet) error
	Delete(id int) error
}