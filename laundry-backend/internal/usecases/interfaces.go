package usecases

import (
	"laundry-backend/internal/entities"
)

type AuthUsecase interface {
	Login(request entities.LoginRequest) (*entities.LoginResponse, error)
}

type BrandUsecase interface {
	CreateBrand(request entities.RegisterBrandRequest) error
	GetBrandByID(id int) (*entities.Brand, error)
	GetAllBrands() ([]entities.Brand, error)
	GetAllBrandsDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error)
	UpdateBrand(id int, request entities.RegisterBrandRequest) error
	DeleteBrand(id int) error
}

type CabangUsecase interface {
	CreateCabang(request entities.RegisterCabangRequest) error
	GetCabangByID(id int) (*entities.Cabang, error)
	GetCabangsByBrandID(brandID int) ([]entities.Cabang, error)
	GetAllCabangs() ([]entities.Cabang, error)
	GetAllCabangsDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error)
	UpdateCabang(id int, request entities.RegisterCabangRequest) error
	DeleteCabang(id int) error
}

type OutletUsecase interface {
	CreateOutlet(request entities.RegisterOutletRequest) error
	GetOutletByID(id int) (*entities.Outlet, error)
	GetOutletsByCabangID(cabangID int) ([]entities.Outlet, error)
	GetAllOutlets() ([]entities.Outlet, error)
	GetAllOutletsDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error)
	UpdateOutlet(id int, request entities.RegisterOutletRequest) error
	DeleteOutlet(id int) error
}