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

type InquiryRepository interface {
	ValidateServicePackage(id int) (bool, error)
	ValidateEmployee(id int) (bool, error)
	ValidateCustomer(id int) (bool, error)
	InsertTransaction(transaction *entities.Transaction) error
	InsertTransactionDetail(detail *entities.TransactionDetail) error
	GetServicePackagePrice(id int) (float64, error)
}

type EmployeeRepository interface {
	Create(employee *entities.Employee) error
	FindByID(id int) (*entities.Employee, error)
	FindAll() ([]entities.Employee, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Employee, int, int, error)
	Update(employee *entities.Employee) error
	Delete(id int) error
}

type CustomerRepository interface {
	Create(customer *entities.Customer) error
	FindByID(id int) (*entities.Customer, error)
	FindByOutletID(outletID int) ([]entities.Customer, error)
	FindAll() ([]entities.Customer, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Customer, int, int, error)
	Update(customer *entities.Customer) error
	Delete(id int) error
}

type ServiceRepository interface {
	Create(service *entities.Service) error
	FindByID(id int) (*entities.Service, error)
	FindAll() ([]entities.Service, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Service, int, error)
	Update(service *entities.Service) error
	Delete(id int) error
	FindByCategoryID(categoryID int) ([]entities.Service, error)
}

type ServiceCategoryRepository interface {
	Create(category *entities.ServiceCategory) error
	FindByID(id int) (*entities.ServiceCategory, error)
	FindAll() ([]entities.ServiceCategory, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.ServiceCategory, int, int, error)
	Update(category *entities.ServiceCategory) error
	Delete(id int) error
}

type EmployeeAccessRepository interface {
	Create(access *entities.EmployeeAccess) error
	FindByID(id int) (*entities.EmployeeAccess, error)
	FindByUsername(username string) (*entities.EmployeeAccess, error)
	FindAll() ([]entities.EmployeeAccess, error)
	FindAllWithPagination(limit, offset int) ([]entities.EmployeeAccess, int, error)
	Update(access *entities.EmployeeAccess) error
	UpdatePassword(id int, password string) error
	UpdateLastLogin(id int) error
	Delete(id int) error
	FindByOutletID(outletID int) ([]entities.EmployeeAccess, error)
	AuthenticateEmployee(username, password string) (*entities.EmployeeAccess, error)
}
